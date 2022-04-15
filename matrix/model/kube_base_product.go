package model

import (
	"database/sql"
	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"time"
)

type deployKubeBaseProduct struct {
	dbhelper.DbTable
}

var DeployKubeBaseProduct = &deployKubeBaseProduct{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_DEPLOY_KUBE_BASE_PRODUCT_LIST},
}

type KubeBaseProduct struct {
	Id            int               `db:"id" json:"id"`
	Pid           int               `db:"pid" json:"pid"`
	ClusterId     int               `db:"clusterId" json:"clusterId"`
	Namespace     string            `db:"namespace" json:"namespace"`
	RelyNamespace string            `db:"rely_namespace" json:"rely_namespace"`
	UpdateTime    dbhelper.NullTime `db:"update_time" json:"update_time"`
	CreateTime    dbhelper.NullTime `db:"create_time" json:"create_time"`
	IsDeleted     int               `db:"is_deleted" json:"is_deleted"`
}

func (l *deployKubeBaseProduct) GetByPidAndClusterIdAndNamespace(pid, clusterId int, namespace string) (KubeBaseProduct, error) {
	info := KubeBaseProduct{}
	err := l.GetWhere(nil, dbhelper.MakeWhereCause().
		Equal("pid", pid).And().
		Equal("clusterId", clusterId).And().
		Equal("namespace", namespace).And().
		Equal("is_deleted", 0), &info)
	return info, err
}

func (l *deployKubeBaseProduct) InsertRecord(info KubeBaseProduct) error {
	err := l.GetWhere(nil, dbhelper.MakeWhereCause().Equal("pid", info.Pid).And().
		Equal("is_deleted", 0).And().Equal("clusterId", info.ClusterId).And().Equal("namespace", info.Namespace), &info)
	if err != nil && err == sql.ErrNoRows {
		_, err = l.InsertWhere(dbhelper.UpdateFields{
			"pid":            info.Pid,
			"clusterId":      info.ClusterId,
			"namespace":      info.Namespace,
			"rely_namespace": info.RelyNamespace,
		})
	} else if err == nil {
		err = l.UpdateWhere(dbhelper.MakeWhereCause().Equal("pid", info.Pid).And().Equal("is_deleted", 0).
			And().Equal("clusterId", info.ClusterId).And().Equal("namespace", info.Namespace), dbhelper.UpdateFields{
			"rely_namespace": info.RelyNamespace,
			"update_time":    time.Now(),
		}, false)
	} else {
		return err
	}
	return err
}
