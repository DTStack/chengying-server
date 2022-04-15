package k8s

import (
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/k8s/impl"
	"github.com/kataras/iris/context"
)

var ClusterResourceAPIRoutes = apibase.Route{
	Path:        "cluster/manage",
	SubRoutes:   []apibase.Route{
		{
			Path:        "namespaces",
			GET:         impl.NameSpaceListStatus,
			Docs:        apibase.Docs{
				Desc:   	"query namespaces situation",
			},
			SubRoutes:   []apibase.Route{
				{
					Path: "{namespace:string}",
					GET: impl.NamespaceStatus,
					Docs: apibase.Docs{
						Desc: "query namespace situation",
					},
				},
			},
		},
		{
			Path:        "namespace",
			SubRoutes:   []apibase.Route{
				{
					Path: 		"save",
					POST: 		impl.NamespaceSave,
					Middlewares: []context.Handler{
						apibase.CheckPermission3,
					},
					Docs: 		apibase.Docs{
						Desc: 		"save data about the client in different namespace that represents different permissions",
					},
				},
				{
					Path: 		"agent/generate",
					POST:       impl.AgentGenerate,
					Docs:       apibase.Docs{
						Desc:       "generate agent mode import yaml",
					},
				},
				{
					Path: 		"ping",
					POST: 		impl.NamespacePing,
					Docs: 		apibase.Docs{
						Desc: 		"test the agent mode if connect",
					},
				},
			},
		},
		{
			Path:		"{namespace:string}",
			SubRoutes: 	[]apibase.Route{
				{
					Path: 		"get",
					GET: 		impl.NamespaceGet,
					Docs: 		apibase.Docs{
						Desc: 		"get namespace client info",
					},
				},
				{
					Path: 		"events",
					GET: 		impl.NamespaceEvent,
					Docs: 		apibase.Docs{
						Desc: 		"get namespace event",
					},
				},
				{
					Path: 		"delete",
					POST: 		impl.NamespaceDelete,
					Middlewares: []context.Handler{
						apibase.CheckPermission3,
					},
					Docs: 		apibase.Docs{
						Desc: 		"delete namespace client info",
					},
					SubRoutes: 	[]apibase.Route{
						{
							Path: 	"confirm",
							GET: 	impl.NamespaceDeleteConfirm,
							Middlewares: []context.Handler{
								apibase.CheckPermission3,
							},
							Docs: 	apibase.Docs{
								Desc: "Confirm whether it can be deleted",
							},
						},
					},
				},
			},
		},
	},
}
