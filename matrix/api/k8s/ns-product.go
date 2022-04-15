package k8s

import (
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/k8s/impl"
)

var NSProductAPIRoutes = apibase.Route{
	Path: "product/manage",
	SubRoutes: []apibase.Route{
		{
			Path: 	"{namespace:string}",
			GET: 	impl.GetParentProducts,
			Docs: 	apibase.Docs{
				Desc: 	"get deployed parent product name in namespace",
			},
			SubRoutes: 	[]apibase.Route{
				{
					Path:	"{parent_product_name:string}",
					GET: 	impl.GetProducts,
					Docs: 	apibase.Docs{
						Desc: 	"get deployed product named in parent product",
					},
					SubRoutes: []apibase.Route{
						{
							Path: 		"{product_name:string}",
							GET: 		impl.GetServiceList,
							SubRoutes: 	[]apibase.Route{
								{
									Path: 		"{service_name:string}",
									GET: 		impl.GetService,
								},
							},
						},
					},
				},
			},
		},
	},
}
