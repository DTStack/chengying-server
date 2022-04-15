package model

import (
	"database/sql"
	"time"
)

type ServiceIpNode struct {
	Id          int64     `db:"id"`
	ClusterId   int       `db:"cluster_id"`
	ProductName string    `db:"product_name"`
	ServiceName string    `db:"service_name"`
	NameSpace   string    `db:"namespace"`
	Ip          string    `db:"ip"`
	NodeId      uint      `db:"node_id"`
	CreatedAt   time.Time `db:"created_at"`
}

func GetServiceNodes(clusterId int, product, service string) ([]ServiceIpNode, error) {
	query := `SELECT * FROM deploy_service_ip_node WHERE product_name = ? AND service_name = ? AND cluster_id = ?`
	nodes := make([]ServiceIpNode, 0)
	return nodes, USE_MYSQL_DB().Select(&nodes, query, product, service, clusterId)
}

func GetServiceIpNode(clusterId int, product, service, ip string) (*ServiceIpNode, error) {
	query := "SELECT * FROM deploy_service_ip_node WHERE product_name = ? AND service_name = ? AND ip = ? AND cluster_id=?"
	node := ServiceIpNode{}
	return &node, USE_MYSQL_DB().Get(&node, query, product, service, ip, clusterId)
}

func (s *ServiceIpNode) Create() error {
	query := "INSERT INTO deploy_service_ip_node (cluster_id, product_name, service_name, ip, node_id) VALUES (?, ?, ?, ?, ?)"
	result, err := USE_MYSQL_DB().Exec(query, s.ClusterId, s.ProductName, s.ServiceName, s.Ip, s.NodeId)
	if err != nil {
		return err
	}

	s.Id, err = result.LastInsertId()
	return err
}

func (ServiceIpNode) DeleteByIp(ip string) error {
	const sqlStr = "DELETE FROM deploy_service_ip_node WHERE  ip = ?"
	_, err := USE_MYSQL_DB().Exec(sqlStr, ip)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceIpNode) Delete() error {
	var err error
	if s.Id > 0 {
		query := "DELETE FROM deploy_service_ip_node WHERE id = ?"
		_, err = USE_MYSQL_DB().Exec(query, s.Id)
	} else {
		query := "DELETE FROM deploy_service_ip_node WHERE product_name = ? AND service_name = ? AND ip = ? AND cluster_id = ? "
		_, err = USE_MYSQL_DB().Exec(query, s.ProductName, s.ServiceName, s.Ip, s.ClusterId)
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	return err
}

func UpdateNodeIdWithNodeId(clusterId int, product, service string, oldNodeId, newNodeId uint) error {
	query := "UPDATE deploy_service_ip_node SET node_id = ? where cluster_id = ? AND product_name = ? AND service_name = ? AND node_id = ? "
	_, err := USE_MYSQL_DB().Exec(query, newNodeId, clusterId, product, service, oldNodeId)
	if err == sql.ErrNoRows {
		err = nil
	}
	return err
}

func DeleteNodeByClusterIdProductService(clusterId int, product, service string) error {
	const sqlStr = "DELETE FROM deploy_service_ip_node WHERE  cluster_id = ? AND product_name= ? AND service_name = ? "
	_, err := USE_MYSQL_DB().Exec(sqlStr, clusterId, product, service)
	if err != nil {
		return err
	}
	return nil
}

