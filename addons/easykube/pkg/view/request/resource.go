package request

import "dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"

type Resource struct {
	Data    []byte            `json:"data"`
	Group   string            `json:"group"`
	Kind    string            `json:"kind"`
	Version string            `json:"version"`
	Action  base.DryRunAction `json:"action"`
}
