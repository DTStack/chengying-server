package apis

import (
	molev1 "dtstack.com/dtstack/easymatrix/addons/operator/pkg/apis/mole/v1"
	workloadv1beta1 "dtstack.com/dtstack/easymatrix/addons/operator/pkg/apis/workload/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

var localSchemeBuilder = runtime.SchemeBuilder{
	workloadv1beta1.AddToScheme,
	molev1.AddToScheme,
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToScheme = localSchemeBuilder.AddToScheme
