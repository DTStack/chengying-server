package api

import (
	"dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/impl"
)

var RoleOperationEasyMatrixAPIRoutes = apibase.Route{
	Path: "role",
	SubRoutes: []apibase.Route{{
		Path: "{role_id:int}",
		SubRoutes: []apibase.Route{{
			//http:://xxxx/api/v2/role/{role_id}/permissions
			Path: "permissions",
			GET:  impl.GetRolePermissions,
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "通过role_id获取角色的所有权限点",
				},
			},
		}},
	}, {
		Path: "list",
		GET:  impl.GetRoleList,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "获取角色列表",
			},
		},
	}, {
		//http:://xxxx/api/v2/role/codes
		Path: "codes",
		GET:  impl.GetRolePermissionCodes,
		Docs: apibase.Docs{
			GET: &apibase.ApiDoc{
				Name: "通过role_id获取角色的所有权限code，方便前端比较",
			},
		},
	}},
}
