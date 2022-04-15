package model

import (
	"fmt"
	"strings"
)

func MergeAnnotations(requested map[string]string, existing map[string]string) map[string]string {
	//if existing == nil {
	//	return requested
	//}
	//
	//for k, v := range requested {
	//	existing[k] = v
	//}
	return existing
}

func BuildResourceName(resourceType, parentProductName, productName, serviceName string) string {
	return fmt.Sprintf("%v-%v-%v-%v", resourceType, ConvertDNSRuleName(parentProductName), ConvertDNSRuleName(productName), ConvertDNSRuleName(serviceName))
}

func BuildResourceLabel(parentProductName, productName, serviceName string) string {
	return fmt.Sprintf("%v-%v-%v", ConvertDNSRuleName(parentProductName), ConvertDNSRuleName(productName), ConvertDNSRuleName(serviceName))
}

func BuildPortName(serviceName string, index int) string {
	return fmt.Sprintf("%v-%v", ConvertDNSRuleName(serviceName), index)
}

func ConvertDNSRuleName(s string) string {
	s = strings.Replace(s, "_", "", -1)
	s = strings.ToLower(s)
	return s
}
