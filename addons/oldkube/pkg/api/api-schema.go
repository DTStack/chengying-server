package api

import "dtstack.com/dtstack/easymatrix/go-common/api-base"

var ApiV2Schema = apibase.Route{
	Path: "/api/v1",
	SubRoutes: []apibase.Route{
		ClientGoAPIRoutes,
	},
}
