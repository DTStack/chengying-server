package internal

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconciler interface {
	Reconcile(ctx context.Context, c client.Client, recorder record.EventRecorder) error
}

type MutateFunction func(object runtime.Object) (bool, error)
