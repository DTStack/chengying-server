package resourcequota

import (
	"context"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/kube"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GVK = schema.GroupVersionKind{
	Group:   "",
	Version: "v1",
	Kind:    "ResourceQuota",
}
func New() *corev1.ResourceQuota{
	return &corev1.ResourceQuota{}
}

func Get(ctx context.Context,client kube.Client, namespace string) (*corev1.ResourceQuota,error){
	if client == nil{
		return nil,fmt.Errorf("the namespace client is not exist")
	}
	quotas := &corev1.ResourceQuotaList{}
	if err := client.List(ctx,quotas,namespace);err != nil{
		return nil,err
	}
	if len(quotas.Items) == 0{
		return nil,nil
	}
	quota := quotas.Items[0]
	return &quota,nil
}

func Ping(client kube.Client, namespace string) error{
	ping := &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name: "dtstack-dryru",
		},
	}
	if _,err := client.Get(context.Background(),ping);err != nil{
		return err
	}
	return nil
}

