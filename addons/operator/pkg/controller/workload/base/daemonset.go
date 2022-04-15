package base

import (
	workloadv1beta1 "dtstack.com/dtstack/easymatrix/addons/operator/pkg/apis/workload/v1beta1"
	"dtstack.com/dtstack/easymatrix/addons/operator/pkg/controller/internal"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type DaemonSet struct {
}

func (ds DaemonSet) GroupVersionKind() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "apps",
		Version: "v1",
		Kind:    "DaemonSet",
	}
}

func (ds *DaemonSet) Status(info *CtrlInfo, owner *workloadv1beta1.WorkLoad, partName string) {

}

func (ds *DaemonSet) DefaultObject(partName, ownerName, namespace string) runtime.Object {
	return nil
}

func (ds *DaemonSet) GetMutateFunction(info *CtrlInfo, owner *workloadv1beta1.WorkLoad, part workloadv1beta1.WorkLoadPart) internal.MutateFunction {
	return nil
}
