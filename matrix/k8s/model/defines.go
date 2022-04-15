package model

import (
	"fmt"
	"strings"
)

var (
	SECRET_VERSION  = "v1"
	SECRET_KIND     = "Secret"
	SECRET_GROUP    = ""
	SECRET_RESOURCE = "secrets"

	CONFIGMAP_KIND     = "ConfigMap"
	CONFIGMAP_GROUP    = ""
	CONFIGMAP_VERSION  = "v1"
	CONFIGMAP_RESOURCE = "configmaps"

	MOLE_VERSION  = "v1"
	MOLE_KIND     = "Mole"
	MOLE_GROUP    = "operator.dtstack.com"
	MOLE_RESOURCE = "moles"

	NAMESPACE_VERSION  = "v1"
	NAMESPACE_GROUP    = ""
	NAMESPACE_KIND     = "Namespace"
	NAMESPACELIST_KIND = "NamespaceList"

	NAMESPACE_PREFIX = "dtstack-"

	SupportResources = map[string]struct{}{
		"cpu":    {},
		"memory": {},
	}
)

func BuildResourceName(resourceType, parentProductName, productName, serviceName string) string {
	return fmt.Sprintf("%v-%v-%v-%v", resourceType, ConvertDNSRuleName(parentProductName), ConvertDNSRuleName(productName), ConvertDNSRuleName(serviceName))
}
func ConvertDNSRuleName(s string) string {
	s = strings.Replace(s, "_", "", -1)
	s = strings.ToLower(s)
	return s
}

func BuildResourceNameWithNamespace(resourceType, parentProductName, productName, serviceName, namespace string) string {
	return fmt.Sprintf("%v-%v-%v-%v.%v", resourceType, ConvertDNSRuleName(parentProductName), ConvertDNSRuleName(productName), ConvertDNSRuleName(serviceName), namespace)
}
