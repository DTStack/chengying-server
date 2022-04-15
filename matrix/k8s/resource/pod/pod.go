package pod

import "k8s.io/apimachinery/pkg/runtime/schema"

var GVK = schema.GroupVersionKind{
	Group:   "",
	Version: "v1",
	Kind:    "Pod",
}
