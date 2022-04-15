package api

import (
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/api/impl"
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
)

var ResourceApi = apibase.Route{
	Path: "kube/resource",
	SubRoutes: []apibase.Route{{
		//http:://xxxx/api/v1/kube/resource/get
		Path: "get",
		POST: impl.Get,
		Docs: apibase.Docs{
			Desc: "get k8s resource",
		},
	}, {
		//http:://xxxx/api/v1/kube/resource/delete
		Path: "delete",
		POST: impl.Delete,
		Docs: apibase.Docs{
			Desc: "delete k8s resource",
		},
	}, {
		//http:://xxxx/api/v1/kube/resource/apply
		Path: "apply",
		POST: impl.Apply,
		Docs: apibase.Docs{
			Desc: "apply k8s resource",
		},
	}, {
		//http:://xxxx/api/v1/kube/resource/list
		Path: "list",
		POST: impl.List,
		Docs: apibase.Docs{
			Desc: "list k8s resource",
		},
	}, {
		//http:://xxxx/api/v1/kube/resource/create
		Path: "create",
		POST: impl.Create,
		Docs: apibase.Docs{
			Desc: "create k8s resource",
		},
	}, {
		Path: "dryrun",
		POST: impl.DryRun,
		Docs: apibase.Docs{
			Desc: "dry run",
		},
	}, {
		Path: "update",
		POST: impl.Update,
		Docs: apibase.Docs{
			Desc: "update resource",
		},
	}, {
		Path: "events",
		GET:  impl.Events,
		Docs: apibase.Docs{
			Desc: "get listwathc events",
		},
	}, {
		Path: "status",
		POST: impl.Status,
		Docs: apibase.Docs{
			Desc: "update status",
		},
	}},
}
