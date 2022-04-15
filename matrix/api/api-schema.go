package api

import (
	"dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/k8s"
)

var ApiV2Schema = apibase.Route{
	Path: "/api/v2",
	SubRoutes: []apibase.Route{
		SshOperationEasyMatrixAPIRoutes,
		ProductOperationEasyMatrixAPIRoutes,
		InstanceOperationEasyMatrixAPIRoutes,
		InstanceRecordOperationEasyMatrixAPIRoutes,
		ServiceOperationEasyMatrixAPIRoutes,
		GroupOperationEasyMatrixAPIRoutes,
		CommonOperationEasyMatrixAPIRoutes,
		UserOperationEasyMatrixAPIRoutes,
		ClusterEasyMatrixAPIRoutes,
		RoleOperationEasyMatrixAPIRoutes,
		k8s.ClusterResourceAPIRoutes,
		k8s.NSProductAPIRoutes,
		k8s.InstanceApi,
		DashboardOperationEasyMatrixRoutes,
		InspectReportEasyMatrixRoutes,
		TaskOperationEasyMatrixAPIRoutes,
	},
}
