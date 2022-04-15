package cache

import (
	"database/sql"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	configOnce        sync.Once
	Db                *sqlx.DB
	getDB             = func() *sqlx.DB { return Db }
	SysConfig         = sysConfigManage{}
	sysConfigDatalist = SysConfigDataList{
		tableName: "sys_config",
	}
)

func InitSysConfig() {
	configOnce.Do(func() {
		SysConfig = sysConfigManage{
			splitKey: "sys",
			data:     make(map[string]string),

			sysConfig: &sysConfig{
				&PlatFormSecurity{},
			},
		}
		SysConfig.flushSysConfigManage()
	})
}

type sysConfigManage struct {
	splitKey  string
	data      map[string]string
	searchKey string
	*sysConfig
}

type SysConfigDataList struct {
	tableName string
}

func (s *SysConfigDataList) listSysConfigData() ([]SysConfigData, error) {
	list := []SysConfigData{}
	err := Db.Select(&list, "SELECT * FROM "+s.tableName+" where is_delete = 0")
	return list, err
}

func (s *SysConfigDataList) modifySysConfigData(data map[string]string) error {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	for k, v := range data {
		query := fmt.Sprintf("update %s set config_value = ?,update_at = ? where config_name = ? and is_delete = 0", s.tableName)
		row, err := tx.Exec(query, v, time.Now(), k)
		if err != nil {
			tx.Rollback()
			return err
		}
		if i, _ := row.RowsAffected(); i == 0 {
			continue
		}
	}
	tx.Commit()
	return nil
}

type SysConfigData struct {
	Id          int            `db:"id"`
	ConfigName  string         `db:"config_name"`
	ConfigValue sql.NullString `db:"config_value"`
	CreateAt    time.Time      `db:"create_at"`
	UpdateAt    time.Time      `db:"update_at"`
	IsDelete    int            `db:"is_delete"`
}

type sysConfig struct {
	*PlatFormSecurity `sys:"platformsecurity" json:"-"`
}

// type support [string int []string []int]
type PlatFormSecurity struct {
	LoginEncrypt           string `sys:"login_encrypt" json:"login_encrypt"`
	AccountLogoutSleepTime int    `sys:"account_logout_sleep_time" json:"account_logout_sleep_time"`
	ForceResetPassword     int    `sys:"force_reset_password" json:"force_reset_password"`
	AccountLoginLockSwitch int    `sys:"account_login_lock_switch" json:"account_login_lock_switch"`
	AccountLoginLimitError int    `sys:"account_login_limit_error" json:"account_login_limit_error"`
	AccountLoginLockTime   int    `sys:"account_login_lock_time" json:"account_login_lock_time"`
}

// sync to db and flush cache
func (sm *sysConfigManage) UpdatePlatFormSecurity(data PlatFormSecurity) error {
	var (
		sVal    = reflect.ValueOf(data)
		sType   = reflect.TypeOf(data)
		dataMap = make(map[string]string)
		prefix  = "platformsecurity."
	)
	for i := 0; i < sVal.NumField(); i++ {
		k := sType.Field(i).Tag.Get(sm.splitKey)
		val := fmt.Sprintf("%v", sVal.Field(i).Interface())
		dataMap[prefix+k] = fmt.Sprintf("%v", val)
	}
	log.Infof("[sysConfigManage.UpdatePlatFormSecurity] %s", dataMap)
	err := sysConfigDatalist.modifySysConfigData(dataMap)
	if err != nil {
		log.Errorf("[sysConfigManage.modifySysConfigData] %s", err)
		return err
	}
	sm.flushSysConfigManage()
	return nil
}

func (sm *sysConfigManage) GetSysconfig() *sysConfig {
	return sm.sysConfig
}

func (sm *sysConfigManage) add(key string) {
	if sm.searchKey == "" {
		sm.searchKey = key
	} else {
		sm.searchKey = sm.searchKey + "." + key
	}
}

func (sm *sysConfigManage) delLastKey() {
	if sm.searchKey == "" {
		return
	}
	arr := strings.Split(sm.searchKey, ".")
	sm.searchKey = strings.Join(arr[0:len(arr)-1], ".")
}

func (sm *sysConfigManage) traverse(target interface{}) {
	sVal := reflect.ValueOf(target)
	sType := reflect.TypeOf(target)
	if sType.Kind() == reflect.Ptr {
		sVal = sVal.Elem()
		sType = sType.Elem()
	}

	num := sVal.NumField()
	for i := 0; i < num; i++ {
		k := sType.Field(i).Tag.Get(sm.splitKey)
		if k != "" {
			sm.add(k)
		}
		//判断字段是否为结构体类型，或者是否为指向结构体的指针类型
		if sVal.Field(i).Kind() == reflect.Struct || (sVal.Field(i).Kind() == reflect.Ptr && sVal.Field(i).Elem().Kind() == reflect.Struct) {
			sm.traverse(sVal.Field(i).Interface())
		}
		if k != "" {
			field := sVal.Field(i)
			v, ok := sm.data[sm.searchKey]
			if field.IsValid() && ok {
				kind := field.Type().Kind()
				switch kind {
				case reflect.String:
					sVal.Field(i).Set(reflect.ValueOf(v))
				case reflect.Int:
					if v, err := strconv.Atoi(v); err == nil {
						field.Set(reflect.ValueOf(v))
					}
				case reflect.Slice:
					if field.Type().Elem().Kind() == reflect.Int {
						op := []int{}
						for _, s := range strings.Split(strings.Trim(v, "[]"), " ") {
							i, err := strconv.Atoi(s)
							if err != nil {
								log.Errorf("[sysConfigManage.traverse]reflect.Slice.Int %s", err)
							}
							op = append(op, i)
						}
						field.Set(reflect.ValueOf(op))
					}
					if field.Type().Elem().Kind() == reflect.String {
						op := []string{}
						for _, s := range strings.Split(strings.Trim(v, "[]"), " ") {
							op = append(op, s)
						}
						field.Set(reflect.ValueOf(op))
					}
					//v := strings.Split(v, ",")
					//if len(v) > 0 {
					//	field.Set(reflect.ValueOf(v))
					//}
				}
			}
			sm.delLastKey()
		}
	}
}
func (sm *sysConfigManage) flushSysConfigManage() {
	list, err := sysConfigDatalist.listSysConfigData()
	if err != nil && len(list) == 0 {
		log.Errorf("[sysConfigDatalist.listSysConfigData] %s", err)
		return
	}
	for _, data := range list {
		sm.data[data.ConfigName] = data.ConfigValue.String
	}
	sm.traverse(sm.sysConfig)
}
