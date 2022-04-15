package kube

import (
	"database/sql"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"github.com/jmoiron/sqlx"
)
var (
	getImageStoreByIdSql = "select * from deploy_cluster_image_store where id = :id and is_deleted = 0"
	getImageStoreByIdSts *sqlx.NamedStmt
	getImageStoreByClusterIdSql = "select * from deploy_cluster_image_store where clusterid = :clusterId and is_deleted = 0"
	getImageStoreByClusterIdSts *sqlx.NamedStmt
	DeployClusterImageStore = &deployClusterImageStore{
		PrepareFunc: prepareDeployClusterImageStore,
	}
)
type deployClusterImageStore struct {
	PrepareFunc
}

func prepareDeployClusterImageStore() error{
	var err error
	getImageStoreByIdSts,err = model.USE_MYSQL_DB().PrepareNamed(getImageStoreByIdSql)
	if err != nil{
		log.Errorf("[kube deploy_cluster_image_store]: init sql: %s , error %v",getImageStoreByIdSql,err)
		return err
	}
	getImageStoreByClusterIdSts,err = model.USE_MYSQL_DB().PrepareNamed(getImageStoreByClusterIdSql)
	if err != nil {
		log.Errorf("[kube deploy_cluster_image_store]: init sql: %s , error %v",getImageStoreByClusterIdSql,err)
		return err
	}
	return nil
}

type DeployClusterImageStoreSchema struct {
	Id         int               `db:"id"`
	ClusterId  int               `db:"clusterId"`
	IsDefault  int               `db:"is_default"`
	Name       string            `db:"name"`
	Alias      string            `db:"alias"`
	Address    string            `db:"address"`
	Username   string            `db:"username"`
	Password   string            `db:"password"`
	Email      string            `db:"email"`
	UpdateTime sql.NullTime      `db:"update_time"`
	CreateTime sql.NullTime      `db:"create_time"`
	IsDeleted  int               `db:"is_deleted"`
}

func (s *deployClusterImageStore)GetById(id int) (*DeployClusterImageStoreSchema,error){
	sc := &DeployClusterImageStoreSchema{
		Id:         id,
	}
	if err := getImageStoreByIdSts.Get(sc,sc);err!=nil{
		if err == sql.ErrNoRows{
			return nil,nil
		}
		log.Errorf("[kube deploy_cluster_image_store]: get imagestore %s by id %d error: %v",getImageStoreByIdSql,id,err)
		return nil,err
	}
	return sc,nil
}

func (s *deployClusterImageStore)GetByClusterId(cid int) (*[]DeployClusterImageStoreSchema,error){

	arg := &DeployClusterImageStoreSchema{
		ClusterId: cid,
	}

	rows, err := getImageStoreByClusterIdSts.Queryx(arg)
	if err != nil {
		log.Errorf("[kube deploy_cluster_image_store]: init sql: %s, value %+v , error %v",getImageStoreByClusterIdSql,*arg,err)
		return nil,err
	}
	result := []DeployClusterImageStoreSchema{}
	for rows.Next(){
		imageStore:= DeployClusterImageStoreSchema{}
		if err = rows.StructScan(&imageStore);err != nil{
			log.Errorf("[kube deploy_cluster_image_store]: init sql: %s, error %v",getImageStoreByClusterIdSql,err)
			return nil,err
		}
		result = append(result,imageStore)
	}

	return &result,nil
}