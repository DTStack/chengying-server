package model

import (
	dbhelper "dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"fmt"
)

type autoTest struct {
	dbhelper.DbTable
}

var AutoTest = &autoTest{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_AUTO_TEST},
}

type AutoTestInfo struct {
	ID           int               `db:"id"`
	ClusterId    int               `db:"cluster_id"`
	ProductName  string            `db:"product_name"`
	OperationId  string            `db:"operation_id"`
	TestScript   string            `db:"test_script"`
	ExecStatus   int               `db:"exec_status"`
	ReportUrl    string            `db:"report_url"`
	ErrorMessage string            `db:"error_message"`
	CreateTime   dbhelper.NullTime `db:"create_time"`
	EndTime      dbhelper.NullTime `db:"end_time"`
}

func (a *autoTest) Insert(one AutoTestInfo) error {
	sql := fmt.Sprintf("INSERT INTO %s (cluster_id,product_name,operation_id,test_script,exec_status) VALUES (?,?,?,?,?)", TBL_AUTO_TEST)
	_, err := a.GetDB().Exec(sql, one.ClusterId, one.ProductName, one.OperationId, one.TestScript, one.ExecStatus)
	if err != nil {
		return err
	}
	return nil
}

func (a *autoTest) GetByClusterIdAndProductName(clusterId int, productName string) (*AutoTestInfo, error) {
	var one AutoTestInfo
	sql := fmt.Sprintf("select * from %s where cluster_id = ? and product_name = ?  order by create_time desc limit 1", TBL_AUTO_TEST)
	err := a.GetDB().Get(&one, sql, clusterId, productName)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

func (a *autoTest) UpdateStatusByOperationId(operationId, reportUrl, errorMessage string, execStatus int, endTime dbhelper.NullTime) error {
	sql := fmt.Sprintf("update %s set exec_status =?,report_url=?,error_message=?,end_time=?  where operation_id =?", TBL_AUTO_TEST)
	_, err := a.GetDB().Exec(sql, execStatus, reportUrl, errorMessage, endTime, operationId)
	if err != nil {
		return err
	}
	return nil
}
