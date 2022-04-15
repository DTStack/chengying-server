package model

import (
	"dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"dtstack.com/dtstack/easymatrix/go-common/utils"
)

type strategyAssignList struct {
	dbhelper.DbTable
}

var StrategyAssignList = &strategyAssignList{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_DEPLOY_STRATEGY_ASSIGN_LIST},
}

type StrategyAssignInfo struct {
	ID                int               `db:"id"`
	StrategyId        int               `db:"strategy_id"`
	ProductName       string            `db:"product_name"`
	ParentProductName string            `db:"parent_product_name"`
	ServiceName       string            `db:"service_name"`
	Host              string            `db:"host"`
	IsDeleted         int               `db:"is_deleted"`
	GmtCreate         dbhelper.NullTime `db:"gmt_create"`
	GmtModify         dbhelper.NullTime `db:"gmt_modified"`
}

var _getStrategyAssignListFields = utils.GetTagValues(StrategyAssignInfo{}, "db")

func (l *strategyAssignList) GetStrategyAssignList() (error, []*StrategyAssignInfo) {
	list := []*StrategyAssignInfo{}
	whereCause := dbhelper.WhereCause{}
	whereCause = whereCause.Equal("is_deleted", 0)
	rows, _, err := l.SelectWhere(_getStrategyAssignListFields, whereCause, nil)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	for rows.Next() {
		info := &StrategyAssignInfo{}
		err = rows.StructScan(&info)
		if err != nil {
			return err, nil
		}
		list = append(list, info)
	}
	return nil, list
}
