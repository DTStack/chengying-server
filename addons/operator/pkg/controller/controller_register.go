package controller

import (
	"dtstack.com/dtstack/easymatrix/addons/operator/pkg/controller/mole"
	"dtstack.com/dtstack/easymatrix/addons/operator/pkg/controller/workload"
	"dtstack.com/dtstack/easymatrix/addons/operator/pkg/controller/workloadprocess"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs = []func(manager.Manager) error{
	mole.Add,
	workload.Add,
	workloadprocess.Add,
}

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager) error {
	for _, f := range AddToManagerFuncs {
		if err := f(m); err != nil {
			return err
		}
	}
	return nil
}
