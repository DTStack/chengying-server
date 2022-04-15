package deploy

import (
	"context"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/kube"
	k8sModel "dtstack.com/dtstack/easymatrix/matrix/k8s/model"
	sqlModel "dtstack.com/dtstack/easymatrix/matrix/model"
	"encoding/json"
	kschema "k8s.io/apimachinery/pkg/runtime/schema"
)

func ApplyImageSecret(cache kube.ClientCache,clusterId int, namespace string, store sqlModel.ImageStore) error {
	secret := k8sModel.NewDockerConfigSecret(namespace, store)
	gvr := &kschema.GroupVersionResource{
		Group:    k8sModel.SECRET_GROUP,
		Version:  k8sModel.SECRET_VERSION,
		Resource: k8sModel.SECRET_RESOURCE,
	}
	secretBytes, err := json.Marshal(secret)
	if err != nil {
		return err
	}
	secretDynamic := NewDynamic(secretBytes, gvr, k8sModel.SECRET_KIND)
	if cache == nil{
		return ApplyDynamicResource(secretDynamic, clusterId)
	}
	return cache.GetClient(namespace).Apply(context.Background(),secret)
}
