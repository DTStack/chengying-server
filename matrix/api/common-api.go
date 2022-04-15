package api

import (
	"dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/impl"
)

var CommonOperationEasyMatrixAPIRoutes = apibase.Route{
	Path: "common",
	SubRoutes: []apibase.Route{
		{
			Path: "file2text",
			POST: impl.File2text,
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "upload file convert to text[EasyMatrix API]",
				},
			},
		}, {
			Path: "safetyAudit",
			SubRoutes: []apibase.Route{
				{
					Path: "list",
					GET:  impl.GetSafetyAuditList,
					Docs: apibase.Docs{
						GET: &apibase.ApiDoc{
							Name: "get safety audit list[EasyMatrix API]",
						},
					},
				}, {
					Path: "module",
					GET:  impl.GetSafetyAuditModule,
					Docs: apibase.Docs{
						GET: &apibase.ApiDoc{
							Name: "get safety audit module[EasyMatrix API]",
						},
					},
				}, {
					Path: "operation",
					GET:  impl.GetSafetyAuditOperation,
					Docs: apibase.Docs{
						GET: &apibase.ApiDoc{
							Name: "get safety audit operation[EasyMatrix API]",
						},
					},
				},
			},
		},
	},
}
