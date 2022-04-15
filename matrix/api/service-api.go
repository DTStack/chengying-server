package api

import (
	"dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/impl"
	"github.com/kataras/iris/context"
)

var ServiceOperationEasyMatrixAPIRoutes = apibase.Route{
	Path: "service",
	SubRoutes: []apibase.Route{
		{
			Path: "{pid:int}",
			SubRoutes: []apibase.Route{
				{
					Path: "{service_name:string}",
					SubRoutes: []apibase.Route{{
						Path: "start",
						Middlewares: []context.Handler{
							apibase.CheckPermission3,
						},
						POST: impl.ServiceStart,
						Docs: apibase.Docs{
							GET: &apibase.ApiDoc{
								Name: "service start[EasyMatrix API]",
							},
						},
					}, {
						Path: "stop",
						Middlewares: []context.Handler{
							apibase.CheckPermission3,
						},
						POST: impl.ServiceStop,
						Docs: apibase.Docs{
							GET: &apibase.ApiDoc{
								Name: "service stop[EasyMatrix API]",
							},
						},
					}, {
						Path: "rolling_restart",
						Middlewares: []context.Handler{
							apibase.CheckPermission3,
						},
						POST: impl.ServiceRollingRestart,
						Docs: apibase.Docs{
							GET: &apibase.ApiDoc{
								Name: "service rolling restart[EasyMatrix API]",
							},
						},
					}, {
						Path: "config_update",
						Middlewares: []context.Handler{
							apibase.CheckPermission1,
						},
						POST: impl.ServiceRollingConfigUpdate,
						Docs: apibase.Docs{
							GET: &apibase.ApiDoc{
								Name: "service config update[EasyMatrix API]",
							},
						},
					}},
				},
			},
		}, {
			Path: "license",
			POST: impl.License,
		},
	},
}
