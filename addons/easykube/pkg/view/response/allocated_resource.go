package response

type AllocatedResponse struct {
	Nodes          int    `json:"nodes"`
	ErrorNodes     int    `json:"error_nodes"`
	MemSizeDisplay string `json:"mem_size_display"`
	MemUsedDisplay string `json:"mem_used_display"`
	CpuSizeDisplay string `json:"cpu_size_display"`
	CpuUsedDisplay string `json:"cpu_used_display"`
	PodSizeDisplay int64  `json:"pod_size_display"`
	PodUsedDisplay int    `json:"pod_used_display"`
}
