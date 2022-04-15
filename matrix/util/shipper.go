package util

var apiShipper = &ApiShipperPwd{}

type ApiShipperPwd struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Group    string `json:"group"`
}

type PwdConnectParams struct {
	ApiShipperPwd
	ClusterId   int    `json:"cluster_id"`
	ClusterType string `json:"cluster_type"`
	Role        string `json:"role"`
}

type PwdInstallParams struct {
	ApiShipperPwd
	ClusterId   int    `json:"cluster_id"`
	ClusterType string `json:"cluster_type"`
	Role        string `json:"role"`
	Cmd         string `json:"cmd"`
}

type ApiShipperPk struct {
	Host  string `json:"host"`
	Port  string `json:"port"`
	User  string `json:"user"`
	Pk    string `json:"pk"`
	Group string `json:"group"`
}

type PkConnectParams struct {
	ApiShipperPk
	ClusterId   int    `json:"cluster_id"`
	ClusterType string `json:"cluster_type"`
	Role        string `json:"role"`
}

type PkInstallParams struct {
	ApiShipperPk
	ClusterId   int    `json:"cluster_id"`
	ClusterType string `json:"cluster_type"`
	Role        string `json:"role"`
	Cmd         string `json:"cmd"`
}

type ApiShipperCheck struct {
	Aid int `json:"aid"`
}

type ApiShipperCheckByIp struct {
	Ip string `json:"ip"`
}
