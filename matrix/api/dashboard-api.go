package api

import (
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/impl"
)

var DashboardOperationEasyMatrixRoutes = apibase.Route{
	Path: "dashboard",
	SubRoutes: []apibase.Route{
		{
			Path: "import",
			POST: impl.ImportDashboard,
			Docs: apibase.Docs{
				POST: &apibase.ApiDoc{
					Name: "导入grafana仪表盘接口",
				},
			},
		},
		{
			Path: "export",
			GET: impl.ExportDashboard,
			Docs: apibase.Docs{
				POST: &apibase.ApiDoc{
					Name: "导出grafana仪表盘接口",
				},
			},
		}, {
			Path: "alerts",
			GET:  impl.GetDashboardAlerts,
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "获取告警规则列表",
				},
			},
			SubRoutes: []apibase.Route{
				{
					Path: "pause",
					POST: impl.DashboardAlertsPause,
					Docs: apibase.Docs{
						POST: &apibase.ApiDoc{
							Name: "告警规则停止开启",
						},
					},
				},
			},
		},
	},
}
