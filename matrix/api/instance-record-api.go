package api

import (
	"dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/impl"
	"github.com/kataras/iris/context"
)

var InstanceRecordOperationEasyMatrixAPIRoutes = apibase.Route{
	Path: "instance_record",
	SubRoutes: []apibase.Route{
		{
			Path: "{id:int}",
			SubRoutes: []apibase.Route{
				{
					Middlewares: []context.Handler{
						apibase.CheckPermission1,
					},
					Path: "force_stop",
					POST: impl.ForceStop,
					Docs: apibase.Docs{
						GET: &apibase.ApiDoc{
							Name: "force stop instance specified by record id[EasyMatrix API]",
						},
					},
				}, {
					Middlewares: []context.Handler{
						apibase.CheckPermission1,
					},
					Path: "force_uninstall",
					POST: impl.ForceUninstall,
					Docs: apibase.Docs{
						GET: &apibase.ApiDoc{
							Name: "force uninstall instance specified by record id[EasyMatrix API]",
						},
					},
				},
			},
		},
	},
}
