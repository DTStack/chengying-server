package model

import (
	"dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"dtstack.com/dtstack/easymatrix/go-common/utils"
)

type strategyResourcceList struct {
	dbhelper.DbTable
}

var StrategyResourceList = &strategyResourcceList{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_DEPLOY_STRATEGY_RESOURCE_LIST},
}

type StrategyResourceInfo struct {
	ID         int               `db:"id"`
	StrategyId int               `db:"strategy_id"`
	Content    string            `db:"content"`
	IsDeleted  int               `db:"is_deleted"`
	GmtCreate  dbhelper.NullTime `db:"gmt_create"`
	GmtModify  dbhelper.NullTime `db:"gmt_modified"`
}

var _getStrategyResourceListFields = utils.GetTagValues(StrategyResourceInfo{}, "db")

func (l *strategyResourcceList) GetStrategyResourceList() (error, []*StrategyResourceInfo) {
	list := []*StrategyResourceInfo{}
	whereCause := dbhelper.WhereCause{}
	whereCause = whereCause.Equal("is_deleted", 0)
	rows, _, err := l.SelectWhere(_getStrategyResourceListFields, whereCause, nil)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	for rows.Next() {
		info := &StrategyResourceInfo{}
		err = rows.StructScan(&info)
		if err != nil {
			return err, nil
		}
		list = append(list, info)
	}
	return nil, list
}
