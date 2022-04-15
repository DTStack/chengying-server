package kube

import (
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"github.com/jmoiron/sqlx"
)

var (
	getMoudleSql = "select * from import_init_moudle where is_deleted = 0"
	getMoudleSts *sqlx.Stmt
	ImportInitMoudle = &importInitMoudle{
		PrepareFunc: prepareImportInitMoudle,
	}
)
type importInitMoudle struct {
	PrepareFunc
}

func prepareImportInitMoudle() error{
	var err error
	getMoudleSts,err = model.USE_MYSQL_DB().Preparex(getMoudleSql)
	if err != nil{
		log.Errorf("[kube import_init_moudle]: init sql: %s , error %v",getMoudleSql,err)
		return err
	}
	return nil
}
type ImportInitMoudleSchema struct {
	Id				int 	`db:"id"`
	ServiceAccount 	string 	`db:"service_account"`
	Role 			string 	`db:"role"`
	RoleBinding 	string 	`db:"role_binding"`
	Operator 		string 	`db:"operator"`
	LogConfig       string  `db:"log_config"`
	IsDeleted 		int 	`db:"is_deleted"`
}

func (i *importInitMoudle) GetInitMoudle() (*ImportInitMoudleSchema,error){
	sc := &ImportInitMoudleSchema{}
	if err := getMoudleSts.Get(sc);err!=nil{
		log.Errorf("[kube import_init_moudle]: get init moudle error %v",err)
		return nil ,err
	}
	return sc,nil
}
