package model

import (
	"database/sql"
	"time"

	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
)

type uncheckedService struct {
	dbhelper.DbTable
}

var DeployUncheckedService = &uncheckedService{
	dbhelper.DbTable{GetDB: USE_MYSQL_DB, TableName: TBL_DEPLOY_UNCHECKED_SERVICE},
}

type DeployUncheckedServiceInfo struct {
	ID                int               `db:"id"`
	ClusterId         int               `db:"cluster_id"`
	Pid               int               `db:"pid"`
	UncheckedServices string            `db:"unchecked_services"`
	UpdateDate        dbhelper.NullTime `db:"update_time"`
	CreateDate        dbhelper.NullTime `db:"create_time"`
	Namespace         string            `db:"namespace"`
}

func (us *uncheckedService) GetUncheckedServicesByPidClusterId(pid, clusterId int, namespace string) (info *DeployUncheckedServiceInfo, err error) {
	info = &DeployUncheckedServiceInfo{}
	whereCause := dbhelper.WhereCause{}
	whereCause = whereCause.Equal("pid", pid)
	if clusterId > 0 {
		whereCause = whereCause.And()
		whereCause = whereCause.Equal("cluster_id", clusterId)
	}
	if namespace != "" {
		whereCause = whereCause.And()
		whereCause = whereCause.Equal("namespace", namespace)
	}
	if err = us.GetWhere(nil, whereCause, info); err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (us *uncheckedService) SetUncheckedService(pid, clusterId int, uncheckedServices string) error {
	return us.UpdateWhere(dbhelper.MakeWhereCause(), dbhelper.UpdateFields{
		"pid":                pid,
		"cluster_id":         clusterId,
		"unchecked_services": uncheckedServices,
		"update_time":        time.Now(),
	}, true)
}
