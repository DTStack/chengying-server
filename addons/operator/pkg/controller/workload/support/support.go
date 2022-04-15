package support

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clinetgoSchema "k8s.io/client-go/kubernetes/scheme"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var SupportSchema = map[string]schema.GroupVersionKind{
	"deployment": {
		Group:   "apps",
		Version: "v1",
		Kind:    "Deployment",
	},
	"statefulset": {
		Group:   "apps",
		Version: "v1",
		Kind:    "StatefulSet",
	},
	"job": {
		Group:   "batch",
		Version: "v1",
		Kind:    "Job",
	},
	"daemonset": {
		Group:   "apps",
		Version: "v1",
		Kind:    "DaemonSet",
	},
	"conf": {
		Group:   "",
		Version: "v1",
		Kind:    "ConfigMap",
	},
	"service": {
		Group:   "",
		Version: "v1",
		Kind:    "Service",
	},
}

var (
	BoundTypeContainer      = "container"
	BoundTypeInitContainer  = "init-container"
	BoundTypeVolume         = "volume"
	BoundTypePvc            = "pvc"
	CreateTypeService       = "service"
	CreateTypeConfigmap     = "conf"
	WorkloadTypeDeployment  = "deployment"
	WorkloadTypeStatefulset = "statefulset"
	WorkloadTypeJob         = "job"
	WorkloadTypeDaemonset   = "daemonset"
)

var log = logf.Log.WithName("support-schema")

func GetGvk(typ string) (schema.GroupVersionKind, error) {
	gvk, exist := SupportSchema[typ]
	if !exist {
		return schema.GroupVersionKind{}, fmt.Errorf("type %s is not support", typ)
	}
	return gvk, nil
}

func GetTypes() (map[string]runtime.Object, error) {
	supports := make(map[string]runtime.Object, len(SupportSchema))
	for k, gvk := range SupportSchema {
		// schema don't need to regist the crd again, it is registed when the cmd is up in the new contollermanager code
		obj, err := clinetgoSchema.Scheme.New(gvk)
		if err != nil {
			log.Error(err, "get type from schema failed", "type", k)
			return nil, err
		}
		supports[k] = obj
	}
	return supports, nil
}
