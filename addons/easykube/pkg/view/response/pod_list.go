package response

type NodePod struct {
	LocalIp     string  `json:"local_ip"`
	PodUsed     int     `json:"pod_used"`
	PodSize     int64   `json:"pod_size"`
	PodUsagePct float64 `json:"pod_usage_pct"`
}

type PodListResponse struct {
	List []NodePod `json:"list"`
}
