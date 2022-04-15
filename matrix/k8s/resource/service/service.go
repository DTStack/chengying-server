package service

import (
	"context"
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/kube"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)
var GVK = schema.GroupVersionKind{
	Group:   "",
	Version: "v1",
	Kind:    "Service",
}

func New() *corev1.Service{
	return &corev1.Service{}
}

func Ping(client kube.Client, namespace string) error{
	ping := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       GVK.Kind,
			APIVersion: GVK.Group+"/"+GVK.Version,
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      "dtstack-dryru",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port: 5555,
				},
			},
		},
		Status: corev1.ServiceStatus{},
	}
	if _,err := client.Get(context.Background(),ping);err != nil{
		return err
	}

	if err := client.DryRun(base.Create,ping);err != nil{
		return err
	}
	return nil
}
