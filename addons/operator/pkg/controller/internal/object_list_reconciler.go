package internal

import (
	"context"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sort"
)

type ObjectListReconciler struct {
	Reconcilers map[int]Reconciler
}

func (list *ObjectListReconciler) Reconcile(ctx context.Context, c client.Client, recorder record.EventRecorder) error {
	reconcilers := list.sort()
	for _, r := range reconcilers {
		if err := r.Reconcile(ctx, c, recorder); err != nil {
			return err
		}
	}
	return nil
}

func (list *ObjectListReconciler) Append(r Reconciler, index int) {
	if list.Reconcilers == nil {
		list.Reconcilers = map[int]Reconciler{}
	}
	list.Reconcilers[index] = r
}

func (list *ObjectListReconciler) sort() []Reconciler {
	keys := make([]int, 0, len(list.Reconcilers))
	for k := range list.Reconcilers {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	reconcilers := make([]Reconciler, 0, len(list.Reconcilers))
	for _, k := range keys {
		reconcilers = append(reconcilers, list.Reconcilers[k])
	}
	return reconcilers
}
