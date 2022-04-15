package crd

import "k8s.io/apimachinery/pkg/runtime/schema"

type GroupVersionResourceKind interface {
	GroupVersionResource() schema.GroupVersionResource
	GroupVersionKind() schema.GroupVersionKind
}

type gvrk struct {
	gv       schema.GroupVersion
	kind     string
	resource string
}

func (g gvrk) GroupVersionResource() schema.GroupVersionResource {
	return g.gv.WithResource(g.resource)
}

func (g gvrk) GroupVersionKind() schema.GroupVersionKind {
	return g.gv.WithKind(g.kind)
}

var MoleGvrk = gvrk{
	gv: schema.GroupVersion{
		Group:   "operator.dtstack.com",
		Version: "v1",
	},
	kind:     "Mole",
	resource: "moles",
}

var WorkloadProcessGvrk = gvrk{
	gv: schema.GroupVersion{
		Group:   "dtstack.com",
		Version: "v1beta1",
	},
	kind:     "WorkloadProcess",
	resource: "workloadprocess",
}
