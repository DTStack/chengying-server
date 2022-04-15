package view

type NamespaceEventRsp struct {
	Size int `json:"size"`
	Events []NamespaceEvent `json:"events"`
}

type NamespaceEvent struct {
	Id  int 	`json:"id"`
	Time string `json:"time"`
	Type string `json:"type"`
	Reason string `json:"reason"`
	Resource string `json:"resource"`
	Message string `json:"message"`
}
