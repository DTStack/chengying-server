package response

type Namespace struct {
	Name string `json:"namespace"`
}

type NamespaceListResponse struct {
	Namespaces []Namespace `json:"namespaces"`
}
