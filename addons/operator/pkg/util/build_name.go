package util

import (
	"fmt"
	"strings"
)

func BuildWorkloadName(productName, serviceName string) string {
	return fmt.Sprintf("%s-%s", strings.ToLower(strings.Replace(productName, "_", "-", -1)), strings.ToLower(strings.Replace(serviceName, "_", "-", -1)))
}

func BuildBaseName(workloadname, partName string) string {
	return fmt.Sprintf("%s-%s",
		strings.ToLower(workloadname),
		strings.ToLower(partName),
	)
}

func BuildStepName(baseName, stepName string) string {
	if len(stepName) == 0 {
		return baseName
	}
	return fmt.Sprintf("%s-%s", strings.ToLower(baseName), strings.ToLower(stepName))
}
