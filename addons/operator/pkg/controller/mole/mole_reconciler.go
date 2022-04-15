package mole

import (
	molev1 "dtstack.com/dtstack/easymatrix/addons/operator/pkg/apis/mole/v1"
	"dtstack.com/dtstack/easymatrix/addons/operator/pkg/controller/common"
	"dtstack.com/dtstack/easymatrix/addons/operator/pkg/controller/model"
)

type MoleReconciler struct {
	DsHash     string
	ConfigHash string
	PluginsEnv string
	Name       string
}

func NewMoleReconciler(name string) *MoleReconciler {
	return &MoleReconciler{
		DsHash:     "",
		ConfigHash: "",
		PluginsEnv: "",
		Name:       name,
	}
}

func (i *MoleReconciler) Reconcile(state *common.ServiceState, cr *molev1.Mole) common.DesiredServiceState {
	desired := common.DesiredServiceState{}
	if cr.Spec.Product.Service[i.Name].IsJob {
		return desired.AddAction(i.jobReconclie(state, cr))
	}
	desired = desired.AddAction(i.getMoleDeploymentDesiredState(state, cr))
	desired = desired.AddAction(i.getMoleServiceDesiredState(state, cr))
	if cr.Spec.Product.Service[i.Name].IsDeployIngress {
		desired = desired.AddAction(i.getMoleIngressDesiredState(state, cr))
	}
	return desired
}

func (i *MoleReconciler) jobReconclie(state *common.ServiceState, cr *molev1.Mole) common.ServiceAction {
	if state.MoleJob == nil {
		return common.GenericCreateAction{
			Ref: model.MoleJob(cr, i.Name),
			Msg: "create Mole Job",
		}
	}
	return nil
}

func (i *MoleReconciler) getMoleServiceDesiredState(state *common.ServiceState, cr *molev1.Mole) common.ServiceAction {
	if state.MoleService == nil {
		return common.GenericCreateAction{
			Ref: model.MoleService(cr, i.Name),
			Msg: "create Mole service",
		}
	}

	return common.GenericUpdateAction{
		Ref: model.MoleServiceReconciled(cr, state.MoleService, i.Name),
		Msg: "update Mole service",
	}
}

func (i *MoleReconciler) getMoleIngressDesiredState(state *common.ServiceState, cr *molev1.Mole) common.ServiceAction {
	if state.MoleIngress == nil {
		return common.GenericCreateAction{
			Ref: model.MoleIngress(cr, i.Name),
			Msg: "create Mole ingress",
		}
	}
	return common.GenericUpdateAction{
		Ref: model.MoleIngressReconciled(cr, state.MoleIngress, i.Name),
		Msg: "update Mole ingress",
	}
}

func (i *MoleReconciler) getMoleDeploymentDesiredState(state *common.ServiceState, cr *molev1.Mole) common.ServiceAction {
	if state.MoleDeployment == nil {
		return common.GenericCreateAction{
			Ref: model.MoleDeployment(cr, i.Name),
			Msg: "create Mole deployment",
		}
	}
	return common.GenericUpdateAction{
		Ref: model.MoleDeploymentReconciled(cr, state.MoleDeployment, i.Name),
		Msg: "update Mole deployment",
	}
}
