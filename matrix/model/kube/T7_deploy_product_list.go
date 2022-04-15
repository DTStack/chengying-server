package kube

import (
	"database/sql"
)

type DeployProductListSchema struct {
	ID                 int               `db:"id"`
	ParentProductName  string            `db:"parent_product_name"`
	ProductName        string            `db:"product_name"`
	ProductNameDisplay string            `db:"product_name_display"`
	ProductVersion     string            `db:"product_version"`
	Product            []byte            `db:"product"`
	ProductParsed      []byte            `db:"product_parsed"`
	IsCurrentVersion   int               `db:"is_current_version"`
	Status             string            `db:"status"`
	DeployUUID         string            `db:"deploy_uuid"`
	AlertRecover       int               `db:"alert_recover"`
	DeployTime         sql.NullTime		 `db:"deploy_time"`
	CreateTime         sql.NullTime		 `db:"create_time"`
	UserId             int               `db:"user_id"`
	Schema             []byte            `db:"schema"`
	ProductType        int               `db:"product_type"`
}
