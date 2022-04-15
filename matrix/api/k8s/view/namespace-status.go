package view

type NamespaceStatusRsp struct {
	Id		  int 		`json:"id"`
	Namespace string 	`json:"namespace"`
	Status    string 	`json:"status"`
	CpuUsed   string 	`json:"cpu_used"`
	CpuTotal  string	`json:"cpu_total"`
	CpuPercent float64		`json:"cpu_percent"`
	MemUsed   string	`json:"memory_used"`
	MemTotal  string	`json:"memory_total"`
	MemPercent float64 		`json:"mem_percent"`
	User 	  string 	`json:"user"`
	UpdateTime string	`json:"update_time"`
	Type 	string 		`json:"type"`
}

