package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	ProcessPending string = "Pending"
	ProcessFinish  string = "finish"
	ProcessFail    string = "fail"
)

type ProcessStatus struct {
	Phase string `json:"phase,omitempty"`
}

// it is used to deploy mutilple services
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=workloadprocess,scope=Namespaced
type WorkloadProcess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkloadProcessSpec `json:"spec"`
	Status            ProcessStatus       `json:"status,omitempty"`
}

type WorkloadProcessSpec struct {
	DeployUUId        string                     `json:"deploy_id"`
	LastDeployUUId    string                     `json:"last_deploy_id",omitempty`
	ProductId         int                        `json:"product_id"`
	ProductName       string                     `json:"product_name",omitempty`
	ProductVersion    string                     `json:"product_version,omitempty"`
	ParentProductName string                     `json:"parent_product_name,omitempty"`
	ClusterId         int                        `json:"cluster_id"`
	WorkLoads         map[string]ServiceWorkload `json:"work_loads,omitempty"`
}

type ServiceWorkload struct {
	Version  string   `json:"version"`
	Group    string   `json:"group"`
	WorkLoad WorkLoad `json:"workload"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type WorkloadProcessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkloadProcess `json:"items"`
}
