package node

type Node struct {
	Aid       int    `json:"aid"`
	ClusterId int    `json:"clusterId"`
	Sid       string `json:"sid"`
	Roles     string `json:"roles"`
	Name      string `json:"name"`
}

type HostNode struct {
	Status  int    `db:"status" json:"status"`
	Steps   int    `db:"steps" json:"steps"`
	Ip      string `db:"ip" json:"ip"`
	Roles   string `db:"roles" json:"roles"`
	Yaml    string `db:"yaml" json:"yaml"`
	Version string `db:"version" json:"version"`
}
