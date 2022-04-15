package view

type InstanceReplicaReq struct {
	Replica     int `json:"replica"`
	ProductName string
	ServiceName string
	Namespace   string `json:"namespace"`
}
