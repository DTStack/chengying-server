package workloadprocess

import (
	workloadv1beta1 "dtstack.com/dtstack/easymatrix/addons/operator/pkg/apis/workload/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GVK = schema.GroupVersionKind{
	Group:   "dtstack.com",
	Version: "v1beta1",
	Kind:    "WorkloadProcess",
}

func New() *workloadv1beta1.WorkloadProcess{
	return &workloadv1beta1.WorkloadProcess{}
}
