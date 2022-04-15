package model

import (
	"database/sql"
	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"time"
)

type deployKubeProductLock struct {
	dbhelper.DbTable
}

var DeployKubeProductLock = &deployKubeProductLock{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_DEPLOY_KUBE_PRODUCT_LOCK},
}

type KubeProductLock struct {
	Id         int               `db:"id" json:"id"`
	Pid        int               `db:"pid" json:"pid"`
	ClusterId  int               `db:"clusterId" json:"clusterId"`
	Namespace  string            `db:"namespace" json:"namespace"`
	IsDeploy   int               `db:"is_deploy" json:"is_deploy"`
	UpdateTime dbhelper.NullTime `db:"update_time" json:"update_time"`
	CreateTime dbhelper.NullTime `db:"create_time" json:"create_time"`
	IsDeleted  int               `db:"is_deleted" json:"is_deleted"`
}

func (l *deployKubeProductLock) GetByPidAndClusterIdAndNamespace(pid, clusterId int, namespace string) (KubeProductLock, error) {
	info := KubeProductLock{}
	err := l.GetWhere(nil, dbhelper.MakeWhereCause().
		Equal("pid", pid).And().
		Equal("clusterId", clusterId).And().
		Equal("namespace", namespace).And().
		Equal("is_deleted", 0), &info)
	return info, err
}

func (l *deployKubeProductLock) InsertOrUpdateRecord(info KubeProductLock) error {
	tmp := KubeProductLock{}
	err := l.GetWhere(nil, dbhelper.MakeWhereCause().Equal("pid", info.Pid).And().
		Equal("is_deleted", 0).And().Equal("clusterId", info.ClusterId).And().Equal("namespace", info.Namespace), &tmp)
	if err != nil && err == sql.ErrNoRows {
		_, err = l.InsertWhere(dbhelper.UpdateFields{
			"pid":       info.Pid,
			"clusterId": info.ClusterId,
			"namespace": info.Namespace,
			"is_deploy": info.IsDeploy,
		})
	} else if err == nil {
		err = l.UpdateWhere(dbhelper.MakeWhereCause().Equal("pid", info.Pid).And().Equal("is_deleted", 0).
			And().Equal("clusterId", info.ClusterId).And().Equal("namespace", info.Namespace), dbhelper.UpdateFields{
			"is_deploy":   info.IsDeploy,
			"update_time": time.Now(),
		}, false)
	} else {
		return err
	}
	return err
}
