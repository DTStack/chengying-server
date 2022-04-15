package k8s

import (
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/k8s/impl"
	"github.com/kataras/iris/context"
)

var InstanceApi = apibase.Route{
	Path:        "instance",
	SubRoutes: 		[]apibase.Route{
		{
			Path: 	"{product_name:string}",
			SubRoutes: []apibase.Route{
				{
					Path:  	"{servce_name:string}",
					SubRoutes: []apibase.Route{
						{
							Path: 	"replica",
							POST: 	impl.InstanceReplica,
							Middlewares: []context.Handler{
								apibase.CheckPermission3,
							},
						},
					},
				},
			},
		},
	},

}
