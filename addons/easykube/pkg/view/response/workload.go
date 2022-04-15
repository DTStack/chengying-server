package response

type WorkLoad struct {
	Load     int `json:"load"`
	Capacity int `json:"capacity"`
}

type WorkLoadResponse struct {
	CornJobs               WorkLoad `json:"CornJobs"`
	Jobs                   WorkLoad `json:"Jobs"`
	Pods                   WorkLoad `json:"Pods"`
	DaemonSets             WorkLoad `json:"DaemonSets"`
	Deployments            WorkLoad `json:"Deployments"`
	ReplicaSets            WorkLoad `json:"ReplicaSets"`
	ReplicationControllers WorkLoad `json:"ReplicationControllers"`
}
