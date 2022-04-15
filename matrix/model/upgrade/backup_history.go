package upgrade

import (
	"fmt"
	"time"

	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"dtstack.com/dtstack/easymatrix/matrix/model"
)

type backupHistory struct {
	dbhelper.DbTable
}

var BackupHistory = &backupHistory{
	dbhelper.DbTable{
		GetDB:     model.USE_MYSQL_DB,
		TableName: model.TBL_BACKUP_HISTORY,
	},
}

type BackupHistoryInfo struct {
	Id          int       `db:"id" json:"id"`
	ClusterId   int       `db:"cluster_id" json:"cluster_id"`
	DbName      string    `db:"db_name" json:"db_name"`
	BackupSql   string    `db:"backup_sql" json:"backup_sql"`
	ProductName string    `db:"product_name" json:"product_name"`
	CreateTime  time.Time `db:"create_time" json:"create_time"`
}

func (b *backupHistory) InsertRecord(clusterId int, dbName, backupSql, productName string) (int64, error) {
	ret, err := b.InsertWhere(dbhelper.UpdateFields{
		"cluster_id":   clusterId,
		"db_name":      dbName,
		"backup_sql":   backupSql,
		"product_name": productName,
	})
	if err != nil {
		return 0, err
	}
	return ret.LastInsertId()
}

func (b *backupHistory) GetLatestRecord(clusterId int, dbName string) (*BackupHistoryInfo, error) {
	info := BackupHistoryInfo{}
	query := fmt.Sprintf("select backup_sql, create_time from %s where cluster_id=? and db_name=? order by create_time "+
		"desc limit 1", b.TableName)
	err := b.Get(&info, query, clusterId, dbName)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
