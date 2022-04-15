package endpoints

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GVK = schema.GroupVersionKind{
	Group:   "",
	Version: "v1",
	Kind:    "Endpoints",
}

func New() *corev1.Endpoints{
	return &corev1.Endpoints{}
}
