package configmap

import (
	"context"
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/kube"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"encoding/json"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GVK = schema.GroupVersionKind{
	Group:   "",
	Version: "v1",
	Kind:    "ConfigMap",
}

func New() *corev1.ConfigMap{
	return &corev1.ConfigMap{}
}

func Ping(client kube.Client, namespace string) error{
	ping := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name: "dtstack-dryru",
		},
	}
	if _,err := client.Get(context.Background(),ping);err != nil{
		return err
	}

	if err := client.DryRun(base.Create,ping);err != nil{
		return err
	}
	return nil
}

func ToObject(bts []byte) (*corev1.ConfigMap,error){
	r,err := base.Schema.New(GVK)
	if err != nil{
		log.Errorf("[configmap]: new object error: %v",err)
		return nil,err
	}
	err = json.Unmarshal(bts,r)
	if err!= nil{
		log.Errorf("[configmap]: json %s unmarshal error: %v",string(bts),err)
		return nil,err
	}
	configmap := r.(*corev1.ConfigMap)
	return configmap,nil
}
