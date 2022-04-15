package model

import (
	"database/sql"
	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/util"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

/*
 @Author: zhijian
 @Date: 2021/4/15 10:40
 @Description:
*/
type productSelectHistory struct {
	dbhelper.DbTable
}

// ProductSelectHistory 该表主要用于 用于部署历史回显
var ProductSelectHistory = &productSelectHistory{
	dbhelper.DbTable{USE_MYSQL_DB, DEPLOY_PRODUCT_SELECT_HISTORY},
}

type ProductSelect struct {
	ClusterId int    `db:"cluster_id" json:"cluster_id"`
	PidList   string `db:"pid_list" json:"pid_list"` //自动部署查看部署历史回显，需要回显的产品包 id list
}

func (ps *productSelectHistory) GetPidListStrByClusterId(clusterId int) (string, error) {
	var pidListStr string
	sql := fmt.Sprintf("select pid_list from %s where cluster_id = ?", DEPLOY_PRODUCT_SELECT_HISTORY)
	err := ps.Get(&pidListStr, sql, clusterId)
	if err != nil {
		return "", err
	}
	return pidListStr, nil
}

func (ps *productSelectHistory) SetPidListStrByClusterId(pidListStr string, clusterId int) error {
	sql := fmt.Sprintf("insert into %s (cluster_id,pid_list) values (?,?) on duplicate key update pid_list = ?", DEPLOY_PRODUCT_SELECT_HISTORY)
	_, err := ps.GetDB().Exec(sql, clusterId, pidListStr, pidListStr)
	if err != nil {
		return err
	}
	return nil
}

func (ps *productSelectHistory) RemovePidByClusterId(currPid string, clusterId int) error {
	log.Debugf("RemovePidByClusterId")
	pidListStr, err := ps.GetPidListStrByClusterId(clusterId)
	//如果 productSelectHistory 表中没有
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	//如果是数据库异常
	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return err
	}
	//如果没有pidListStr
	if strings.TrimSpace(pidListStr) == "" {
		return nil
	}
	pidList := strings.Split(pidListStr, ",")

	if idx := util.IndexOfString(pidList, currPid); idx != -1 {
		pidList = append(pidList[:idx], pidList[idx+1:]...)
	}

	return ps.SetPidListStrByClusterId(strings.Join(pidList, ","), clusterId)
}

func (ps *productSelectHistory) AddPidByClusterId(currPid string, clusterId int) error {
	if currPid == "0" {
		return nil
	}
	pidListStr, err := ps.GetPidListStrByClusterId(clusterId)
	//如果 productSelectHistory 表中没有
	if errors.Is(err, sql.ErrNoRows) {
		return ps.SetPidListStrByClusterId(currPid, clusterId)
	}
	//如果是数据库异常
	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return err
	}
	//如果没有pidListStr
	if strings.TrimSpace(pidListStr) == "" {
		return ps.SetPidListStrByClusterId(currPid, clusterId)
	}
	pidList := strings.Split(pidListStr, ",")

	//如果没找到了
	if idx := util.IndexOfString(pidList, currPid); idx == -1 {
		pidList = append(pidList, currPid)
	}
	return ps.SetPidListStrByClusterId(strings.Join(pidList, ","), clusterId)
}
