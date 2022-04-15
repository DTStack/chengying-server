package listwatch

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type listWatchFactory struct {
	listWatch cache.ListerWatcher
}

// WithCustomObject create listwatch with custom kuber client & object & namespace & fieldSelector.
func WithClientResourceNsLabelSelector(client kubernetes.Clientset, resource, namespace string, labelsSelector labels.Selector) cache.ListerWatcher {
	optionsModifier := func(options *v1.ListOptions) {
		options.LabelSelector = labelsSelector.String()
	}
	listWatcher := cache.NewFilteredListWatchFromClient(client.CoreV1().RESTClient(), resource, namespace, optionsModifier)
	return listWatcher
}

// WithCustomObject create listwatch with custom kuber client & object & namespace & fieldSelector.
func WithConfigResourceNslabelSelector(getter rest.Interface, master, kubeconfig, resource, namespace string, labelsSelector labels.Selector) (error, cache.ListerWatcher) {
	optionsModifier := func(options *v1.ListOptions) {
		options.LabelSelector = labelsSelector.String()
	}
	listWatcher := cache.NewFilteredListWatchFromClient(getter, resource, namespace, optionsModifier)
	return nil, listWatcher
}
