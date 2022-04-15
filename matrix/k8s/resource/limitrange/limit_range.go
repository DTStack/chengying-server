package limitrange

import (
	"context"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/kube"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GVK = schema.GroupVersionKind{
	Group:   "",
	Version: "v1",
	Kind:    "LimitRange",
}

func Get(ctx context.Context,client kube.Client, namespace string) (*corev1.LimitRange,error){
	if client == nil{
		return nil,fmt.Errorf("the namespace client is not exist")
	}
	limitrange := &corev1.LimitRangeList{}
	if err := client.List(ctx,limitrange,namespace);err != nil{
		return nil,err
	}
	if len(limitrange.Items) == 0{
		return nil,nil
	}
	limirange := limitrange.Items[0]
	return &limirange,nil
}
