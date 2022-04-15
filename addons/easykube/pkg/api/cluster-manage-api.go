package api

import (
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/api/impl"
	"dtstack.com/dtstack/easymatrix/go-common/api-base"
)

// clustermanager api
var ClientGoAPIRoutes = apibase.Route{
	Path: "clientgo",
	SubRoutes: []apibase.Route{{
		//http:://xxxx/api/v1/clientgo/allocated
		Path: "allocated",
		GET:  impl.GetAllocated,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "获取集群mem,cpu,pod等资源总使用情况[EasyKube API]",
			},
		},
	}, {
		//http:://xxxx/api/v1/clientgo/top5
		Path: "top5",
		GET:  impl.GetTop5,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "获取cpu、mem资源使用率前五高的机器[EasyKube API]",
			},
		},
	}, {
		//http:://xxxx/api/v1/clientgo/workload
		Path: "workload",
		GET:  impl.GetWorkLoad,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "获取资源工作负载[EasyKube API]",
			},
		},
	}, {
		//http:://xxxx/api/v1/clientgo/allocatedPodList
		Path: "allocatedPodList",
		GET:  impl.GetAllocatedPodList,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "获取集群每个节点的pod资源总使用情况[EasyKube API]",
			},
		},
	}, {
		//http:://xxxx/api/v1/clientgo/componentStatus
		Path: "componentStatus",
		GET:  impl.GetComponentStatus,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "获取集群的组件状态[EasyKube API]",
			},
		},
	}, {
		//http:://xxxx/api/v1/clientgo/extraInfo
		Path: "extraInfo",
		GET:  impl.GetExtraInfo,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "获取节点的role信息，k8s版本信息[EasyKube API]",
			},
		},
	}},
}
