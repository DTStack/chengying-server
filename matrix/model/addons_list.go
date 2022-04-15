package model

import (
	"dtstack.com/dtstack/easymatrix/go-common/db-helper"
	"dtstack.com/dtstack/easymatrix/matrix/base"
)

type addonsList struct {
	dbhelper.DbTable
}

var AddonsList = &addonsList{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_ADDONS_LIST},
}

type AddonsInfo struct {
	Id         int       `db:"id"`
	Type       string    `db:"type"`
	Desc       string    `db:"desc"`
	Os         string    `db:"os"`
	Version    string    `db:"version"`
	Schema     string    `db:"schema"`
	IsDeleted  int       `db:"isDeleted" json:"-"`
	UpdateDate base.Time `db:"updated" json:"updated"`
	CreateDate base.Time `db:"created" json:"created"`
}

func (l *addonsList) GetAddonInfoById(aid string) (error, *AddonsInfo) {
	whereCause := dbhelper.WhereCause{}
	info := AddonsInfo{}
	err := l.GetWhere(nil, whereCause.Equal("id", aid), &info)
	if err != nil {
		return err, &info
	}
	return nil, &info
}
