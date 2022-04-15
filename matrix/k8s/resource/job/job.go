package job

import (
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GVK = schema.GroupVersionKind{
	Group:   "batch",
	Version: "v1",
	Kind:    "Job",
}

func New() *batchv1.Job{
	return &batchv1.Job{}
}
