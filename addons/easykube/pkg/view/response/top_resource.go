package response

type Top5Response struct {
	CpuTop5 []Top5Attribute `json:"cpu_top5"`
	MemTop5 []Top5Attribute `json:"mem_top5"`
}

type Top5Attribute struct {
	Ip    string  `json:"ip"`
	Usage float64 `json:"usage"`
}
