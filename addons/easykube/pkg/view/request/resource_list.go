package request

type ResourceList struct {
	Namespace string `json:"namespace"`
	Group     string `json:"group"`
	Kind      string `json:"kind"`
	Version   string `json:"version"`
}
