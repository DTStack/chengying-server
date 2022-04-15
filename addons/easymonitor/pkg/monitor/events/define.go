package events

import (
	"fmt"
)

var (
	DEFAULT_RESOURCES_POD_KIND        = "Pod"
	DEFAULT_RESOURCES_SERVICE_KIND    = "Service"
	DEFAULT_RESOURCES_DEPLOYMENT_KIND = "Deployment"
	DEFAULT_RESOURCES_INGRESSES_KIND  = "Ingress"
	DEFAULT_RESOURCES_EVENT_KIND      = "Event"
)

type Event struct {
	Namespace   string      `json:"namespace"`
	Resource    string      `json:"resource"`
	Key         string      `json:"key"`
	Operation   string      `json:"operation"`
	Object      interface{} `json:"object"`
	Workspaceid int         `json:"workspaceid"`
}

type EventResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (e *Event) Info() string {
	return fmt.Sprintf("%s-%s-%s-%s", e.Namespace, e.Resource, e.Key, e.Operation)
}
