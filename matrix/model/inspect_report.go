package model

import (
	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"time"
)

type inspectReport struct {
	dbhelper.DbTable
}

var InspectReport = &inspectReport{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_INSPECT_REPORT},
}

type InspectReportInfo struct {
	Id         int               `db:"id" json:"id"`
	Name       string            `db:"name" json:"name"`
	Status     string            `db:"status" json:"status"`
	Progress   int               `db:"progress" json:"progress"`
	CreateTime dbhelper.NullTime `db:"create_time" json:"create_time"`
	UpdateTime dbhelper.NullTime `db:"update_time" json:"update_time"`
	IsDeleted  int               `db:"is_deleted" json:"is_deleted"`
	ClusterId  int               `db:"cluster_id" json:"cluster_id"`
	FilePath   string            `db:"file_path" json:"file_path"`
}

func (i *inspectReport) NewInspectReport(name, status string, clusterId int) (int64, error) {
	result, err := i.InsertWhere(dbhelper.UpdateFields{
		"name":        name,
		"status":      status,
		"progress":    0,
		"create_time": time.Now(),
		"update_time": time.Now(),
		"is_deleted":  0,
		"cluster_id":  clusterId,
		"file_path":   "",
	})
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (i *inspectReport) UpdateProgress(id, progress int, path, status string) error {
	updateFields := dbhelper.UpdateFields{
		"progress":    progress,
		"update_time": time.Now(),
	}
	if path != "" {
		updateFields["file_path"] = path
	}
	if status != "" {
		updateFields["status"] = status
	}
	err := i.UpdateWhere(dbhelper.MakeWhereCause().Equal("id", id), updateFields, false)
	return err
}

func (i *inspectReport) GetById(id int) (InspectReportInfo, error) {
	var info InspectReportInfo
	if err := i.GetWhere(nil, dbhelper.MakeWhereCause().Equal("id", id), &info); err != nil {
		return info, err
	}
	return info, nil
}
