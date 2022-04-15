package api

import (
	"dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/impl"
)

var GroupOperationEasyMatrixAPIRoutes = apibase.Route{
	Path: "group",
	SubRoutes: []apibase.Route{
		{
			Path: "{pid:int}",
			SubRoutes: []apibase.Route{
				{
					Path: "{group_name:string}",
					SubRoutes: []apibase.Route{{
						Path: "start",
						POST: impl.GroupStart,
						Docs: apibase.Docs{
							GET: &apibase.ApiDoc{
								Name: "group start[EasyMatrix API]",
							},
						},
					}, {
						Path: "stop",
						POST: impl.GroupStop,
						Docs: apibase.Docs{
							GET: &apibase.ApiDoc{
								Name: "group stop[EasyMatrix API]",
							},
						},
					}},
				},
			},
		},
	},
}
