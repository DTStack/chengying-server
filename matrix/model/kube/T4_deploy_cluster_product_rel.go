package kube

import (
	"database/sql"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"github.com/jmoiron/sqlx"
)

var (
	selectNamespacedDeployedProductSql = "select * from deploy_cluster_product_rel where clusterId = :clusterId and namespace = :namespace and status = 'deployed'"
	selectNamespecedDeployedProductsts *sqlx.NamedStmt
	DeployClusterProductRel = deployClusterProductRel{
		PrepareFunc: prepareDeplyClusterProductRel,
	}
)

type deployClusterProductRel struct {
	PrepareFunc
}

func prepareDeplyClusterProductRel() error{
	var err error
	selectNamespecedDeployedProductsts,err = model.USE_MYSQL_DB().PrepareNamed(selectNamespacedDeployedProductSql)
	if err != nil{
		log.Errorf("[deploy_cluster_product_rel]: init sql: %s , error %v",selectNamespacedDeployedProductSql,err)
		return err
	}
	return nil
}

type DeployClusterProductRelSchema struct {
	Id            int               `db:"id"`
	Pid           int               `db:"pid"`
	ClusterId     int               `db:"clusterId"`
	Namespace     string            `db:"namespace"`
	ProductParsed []byte            `db:"product_parsed"`
	Status        string            `db:"status"`
	DeployUUID    string            `db:"deploy_uuid"`
	AlertRecover  int               `db:"alert_recover"`
	UserId        int               `db:"user_id"`
	IsDeleted     int               `db:"is_deleted"`
	UpdateTime    sql.NullTime 		`db:"update_time"`
	DeployTime    sql.NullTime 		`db:"deploy_time"`
	CreateTime    sql.NullTime 		`db:"create_time"`
}

func (d *deployClusterProductRel) SelectNamespacedDeployed(clusterid int, namespace string) ([]DeployClusterProductRelSchema,error){
	list := []DeployClusterProductRelSchema{}
	arg := &DeployClusterProductRelSchema{
		ClusterId: clusterid,
		Namespace: namespace,
	}
	rows,err := selectNamespecedDeployedProductsts.Queryx(arg)
	if err != nil && err != sql.ErrNoRows{
		log.Errorf("[deploy_cluster_product_rel]: sql %s, err %v",selectNamespacedDeployedProductSql,err)
		return nil,err
	}
	for rows.Next(){
		tbsc := DeployClusterProductRelSchema{}
		if err = rows.StructScan(&tbsc); err != nil{
			log.Errorf("[deploy_cluster_product_rel]: struct scan to DeployClusterProductRelSchema error: %v",err)
			return nil,err
		}
		list = append(list,tbsc)
	}
	return list,nil
}
