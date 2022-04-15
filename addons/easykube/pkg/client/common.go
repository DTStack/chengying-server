package client

import (
	"context"
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"
	"dtstack.com/dtstack/easymatrix/go-common/log"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	sigClient "sigs.k8s.io/controller-runtime/pkg/client"
)

func GetNodeByIp(ctx context.Context, c *base.Client, hostIp string) (*corev1.Node, error) {
	nodes := &corev1.NodeList{}
	err := c.Lists(ctx, nodes)
	if err != nil {
		return nil, err
	}
	for _, node := range nodes.Items {
		for _, address := range node.Status.Addresses {
			if address.Type == "InternalIP" && address.Address == hostIp {
				return &node, nil
			}
		}
	}
	return nil, fmt.Errorf("cannot match ip %v from cluster", hostIp)
}

func GetNodePods(ctx context.Context, c *base.Client, node *corev1.Node) (*corev1.PodList, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + node.Name +
		",status.phase!=" + string(corev1.PodSucceeded) +
		",status.phase!=" + string(corev1.PodFailed))

	if err != nil {
		log.Errorf("[common] parse selector error : %v", err)
		return nil, err
	}
	opt1 := sigClient.MatchingFieldsSelector{fieldSelector}
	opt2 := sigClient.InNamespace(metav1.NamespaceAll)
	pods := &corev1.PodList{}
	if err := c.Lists(ctx, pods, opt1, opt2); err != nil {
		return nil, err
	}
	return pods, nil
}

func GetNodeConditionStatus(node corev1.Node, conditionType corev1.NodeConditionType) corev1.ConditionStatus {
	for _, condition := range node.Status.Conditions {
		if condition.Type == conditionType {
			return condition.Status
		}
	}
	return corev1.ConditionUnknown
}

func GetNodeConditionMessage(node corev1.Node, conditionType corev1.NodeConditionType) string {
	for _, condition := range node.Status.Conditions {
		if condition.Type == conditionType {
			return condition.Message
		}
	}
	return ""
}
