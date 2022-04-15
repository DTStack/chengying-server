package kube

import (
	"context"
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"
	"k8s.io/apimachinery/pkg/runtime"
)

type Client interface {
	Apply(ctx context.Context, object runtime.Object) error
	Create(ctx context.Context, object runtime.Object) error
	Update(ctx context.Context, object runtime.Object) error
	Delete(ctx context.Context, object runtime.Object) error
	Get(ctx context.Context, object runtime.Object) (bool,error)
	List(ctx context.Context, object runtime.Object, namespace string) error
	Status(ctx context.Context, object runtime.Object) error
    DryRun(action base.DryRunAction,object runtime.Object) error
}

type ClientCache interface {
	Connect(connectStr,workspace string) error
	GetClient(workspace string) Client
	DeleteClient(workspace string)
	Copy() ClientCache
	//need to update?
	//UpdateClient(connectStr,workspace string) error
}

type ImportType string

func (i ImportType) String() string{
	return string(i)
}
var (
	IMPORT_KUBECONFIG              ImportType = "kubeconfig"
	IMPORT_AGENT                   ImportType = "agent"
)
