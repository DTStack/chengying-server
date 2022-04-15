package upgrade

import (
	"database/sql"
	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"time"
)

type upgradeHistory struct {
	dbhelper.DbTable
}

var UpgradeHistory = &upgradeHistory{
	dbhelper.DbTable{
		GetDB:     model.USE_MYSQL_DB,
		TableName: model.TBL_UPGRADE_HISTORY,
	},
}

type HistoryInfo struct {
	Id                int          `db:"id" json:"id"`
	ClusterId         int          `db:"cluster_id" json:"cluster_id"`
	ProductName       string       `db:"product_name" json:"product_name"`
	SourceVersion     string       `db:"source_version" json:"source_version"`
	TargetVersion     string       `db:"target_version" json:"target_version"`
	BackupName        string       `db:"backup_name" json:"backup_name"`
	SourceServiceIp   []byte       `db:"source_service_ip" json:"source_service_ip"`
	SourceConfig      []byte       `db:"source_config" json:"source_config"`
	SourceMultiConfig []byte       `db:"source_multi_config" json:"source_multi_config"`
	CreateTime        sql.NullTime `db:"create_time" json:"create_time"`
	Type              int          `db:"type" json:"type"`
	BackupSql         string       `db:"backup_sql" json:"backup_sql"`
	IsDeleted         bool         `db:"is_deleted" json:"is_deleted"`
}

func (u *upgradeHistory) InsertRecord(clusterId, upgradeType int, productName, sourceVersion, targetVersion, backupName, backupSql string,
	serviceIp, sourceConfig, sourceMultiConfig []byte) (int64, error) {
	var history HistoryInfo
	err := u.GetWhere(nil, dbhelper.MakeWhereCause().Equal("cluster_id", clusterId).And().
		Equal("product_name", productName).And().
		Equal("source_version", sourceVersion).And().
		Equal("target_version", targetVersion).And().
		Equal("backup_name", backupName), &history)
	if err != nil && err == sql.ErrNoRows {
		r, err := u.InsertWhere(dbhelper.UpdateFields{
			"cluster_id":          clusterId,
			"product_name":        productName,
			"source_version":      sourceVersion,
			"target_version":      targetVersion,
			"backup_name":         backupName,
			"source_service_ip":   serviceIp,
			"source_config":       sourceConfig,
			"source_multi_config": sourceMultiConfig,
			"create_time":         time.Now(),
			"type":                upgradeType,
			"backup_sql":          backupSql,
			"is_deleted":          0,
		})
		if err != nil {
			return 0, err
		}
		return r.LastInsertId()
	} else if err == nil {
		return int64(history.Id), nil
	} else {
		return 0, err
	}
}

func (u *upgradeHistory) GetByClsAndProductNameAndSourceVersion(clusterId int, productName, sourceVersion string) ([]HistoryInfo, error) {
	whereClause := dbhelper.MakeWhereCause().Equal("cluster_id", clusterId).And().
		Equal("product_name", productName).And().
		Equal("is_deleted", 0)
	if sourceVersion != "" {
		whereClause = whereClause.And().Equal("source_version", sourceVersion)
	}
	rows, _, err := u.SelectWhere(nil, whereClause, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Errorf("rows close error: %v", err)
			return
		}
	}()
	var infoList []HistoryInfo
	for rows.Next() {
		row := HistoryInfo{}
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		infoList = append(infoList, row)
	}
	return infoList, nil
}

func (u *upgradeHistory) GetOne(clusterId int, productName, sourceVersion, backupName string) (*HistoryInfo, error) {
	whereClause := dbhelper.MakeWhereCause().Equal("cluster_id", clusterId).And().
		Equal("product_name", productName).And().
		Equal("source_version", sourceVersion).And().
		Equal("backup_name", backupName).And().
		Equal("is_deleted", 0)
	var info HistoryInfo
	err := u.GetWhere(nil, whereClause, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
