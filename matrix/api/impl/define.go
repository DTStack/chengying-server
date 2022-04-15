package impl

const (
	COOKIE_CURRENT_CLUSTER_ID    = "em_current_cluster_id"
	COOKIE_INSTALL_CLUSTER_ID    = "em_install_cluster_id"
	COOKIE_PARENT_PRODUCT_NAME   = "em_current_parent_product"
	COOKIE_CURRENT_K8S_NAMESPACE = "em_current_k8s_namespace"
)

type ConfigUpdateParam struct {
	File    string                 `json:"file"`
	Content string                 `json:"content"`
	Values  map[string]interface{} `json:"values"`
	Deleted string                 `json:"deleted"`
}

type AddonInstallParam struct {
	Sid         string                 `json:"sid"`
	AddonId     string                 `json:"addonId"`
	ConfigParam map[string]interface{} `json:"configParam"`
}
