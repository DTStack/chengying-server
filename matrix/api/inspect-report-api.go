package api

import (
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/impl"
	"github.com/kataras/iris/context"
)

var InspectReportEasyMatrixRoutes = apibase.Route{
	Path: "inspect",
	SubRoutes: []apibase.Route{
		{
			Path: "service/status",
			GET:  impl.GetServiceStatus,
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "获取当前集群各产品下服务健康情况",
				},
			},
		}, {
			Path: "alert/history",
			GET:  impl.GetAlertHistory,
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "获取指定时间段内告警历史",
				},
			},
		}, {
			Path: "host/status",
			GET:  impl.GetHostStatus,
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "获取当前集群节点CPU、内存、磁盘健康信息",
				},
			},
		}, {
			Path: "graph",
			SubRoutes: []apibase.Route{
				{
					Path: "config",
					GET:  impl.GetGraphConfig,
					Docs: apibase.Docs{
						GET: &apibase.ApiDoc{
							Name: "获取图表配置列表",
						},
					},
				},
				{
					Path: "data",
					GET:  impl.GetGraphData,
					Docs: apibase.Docs{
						GET: &apibase.ApiDoc{
							Name: "获取图表数据",
						},
					},
				},
			},
		}, {
			Path: "generate",
			POST: impl.StartGenerateReport,
			Docs: apibase.Docs{
				POST: &apibase.ApiDoc{
					Name: "生成巡检报告",
				},
			},
		}, {
			Path: "progress",
			GET:  impl.GetReportProgress,
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "查看生成pdf报告进度",
				},
			},
		}, {
			Path: "download",
			GET:  impl.Download,
			Middlewares: []context.Handler{
				apibase.CheckPermission1,
			},
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "下载巡检报告",
				},
			},
		},
	},
}
