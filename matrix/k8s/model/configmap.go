package model

import (
	"encoding/json"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeConfigMap(namespace, parentProductName, productName, serviceName string, configFiles map[string]string) *v1.ConfigMap {
	//boolTrue := true
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      BuildResourceName("configmap", parentProductName, productName, serviceName),
			Namespace: namespace,
			Labels:    makeLabels(productName, serviceName),
		},
		Data: configFiles,
	}
}

func ConvertConfigMap(data interface{}) (*v1.ConfigMap, error) {
	configBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	conf := &v1.ConfigMap{}
	err = json.Unmarshal(configBytes, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func makeLabels(productName, serviceName string) map[string]string {
	labels := map[string]string{}
	labels["product_name"] = productName
	labels["service_name"] = productName

	return labels
}
