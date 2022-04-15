package model

import (
	"database/sql"
	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"strconv"
	"time"
)

type healthCheck struct {
	dbhelper.DbTable
}

var HealthCheck = &healthCheck{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_SERVICE_HEALTH_CHECK},
}

type HealthCheckInfo struct {
	ID                int               `db:"id"`
	ClusterId         int               `db:"cluster_id"`
	ProductName       string            `db:"product_name"`
	Pid               int               `db:"pid"`
	ServiceName       string            `db:"service_name"`
	AgentId           string            `db:"agent_id"`
	Sid               string            `db:"sid"`
	Ip                string            `db:"ip"`
	ScriptName        string            `db:"script_name"`
	ScriptNameDisplay string            `db:"script_name_display"`
	AutoExec          bool              `db:"auto_exec"`
	Period            string            `db:"period"`
	Retries           int               `db:"retries"`
	ExecStatus        int               `db:"exec_status"`
	ErrorMessage      string            `db:"error_message"`
	StartTime         dbhelper.NullTime `db:"start_time"`
	EndTime           dbhelper.NullTime `db:"end_time"`
	CreateTime        dbhelper.NullTime `db:"create_time"`
	UpdateTime        dbhelper.NullTime `db:"update_time"`
}

func (h *healthCheck) GetInfoByClusterIdAndProductNameAndServiceName(clusterId int, productName, serviceName, hostIp string) ([]HealthCheckInfo, error) {
	s := "select * from " + TBL_SERVICE_HEALTH_CHECK + " where cluster_id = " + "'" + strconv.Itoa(clusterId) + "'" +
		" and product_name = " + "'" + productName + "'" + " and service_name = " + "'" + serviceName + "'"
	if len(hostIp) != 0 {
		s = s + " and ip = '" + hostIp + "'"
	}
	s = s + " order by start_time desc"
	rows, err := USE_MYSQL_DB().Queryx(s)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Errorf("%v", err)
		return nil, err
	}
	list := make([]HealthCheckInfo, 0)
	for rows.Next() {
		info := HealthCheckInfo{}
		if err := rows.StructScan(&info); err != nil {
			log.Errorf("%v", err)
			return nil, err
		}
		list = append(list, info)
	}
	return list, nil
}

func (h *healthCheck) GetInfoById(id int) (HealthCheckInfo, error) {
	info := HealthCheckInfo{}
	err := h.GetWhere(nil, dbhelper.MakeWhereCause().Equal("id", id), &info)
	return info, err
}

func (h *healthCheck) UpdateAutoexecById(id int, autoexec bool) error {
	err := h.UpdateWhere(dbhelper.MakeWhereCause().
		Equal("id", id), dbhelper.UpdateFields{
		"auto_exec":   autoexec,
		"update_time": time.Now(),
	}, false)
	return err
}

func (h *healthCheck) UpdateHealthCheckStatus(id, execStatus int, errorMessage string, startTime, endTime dbhelper.NullTime) error {
	err := h.UpdateWhere(dbhelper.MakeWhereCause().
		Equal("id", id), dbhelper.UpdateFields{
		"exec_status":   execStatus,
		"error_message": errorMessage,
		"start_time":    startTime,
		"end_time":      endTime,
	}, false)
	return err
}
