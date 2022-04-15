package serviceaccount

import (
	"context"
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/kube"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"encoding/json"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GVK = schema.GroupVersionKind{
	Group:   "",
	Version: "v1",
	Kind:    "ServiceAccount",
}

func Convert(obj runtime.Object) *corev1.ServiceAccount{
	return obj.(*corev1.ServiceAccount)
}
func ToObject(bts []byte)(*corev1.ServiceAccount,error){
	r,err := base.Schema.New(GVK)
	if err != nil{
		log.Errorf("[service_account]: new object error: %v",err)
		return nil,err
	}
	err = json.Unmarshal(bts,r)
	if err!= nil{
		log.Errorf("[service_account]: json %s unmarshal error: %v",string(bts),err)
		return nil,err
	}
	sa := r.(*corev1.ServiceAccount)
	return sa,nil
}

func Ping(client kube.Client, namespace string) error{
	sa := &corev1.ServiceAccount{
		ObjectMeta:  metav1.ObjectMeta{
			Namespace: namespace,
			Name: "dtstack-dryrun",
		},
	}
	if _,err := client.Get(context.Background(),sa);err != nil{
		return err
	}

	if err := client.DryRun(base.Create,sa);err != nil{
		return err
	}
	return nil
}
