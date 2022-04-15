package grafana

const (
	GRAFANA_API_ALERTS           = "/api/alerts"
	GRAFANA_API_ALERTS_PAUSE     = "/api/alerts/%s/pause"
	GRAFANA_API_DASHBOARD_IMPORT = "/api/dashboards/import"
	GRAFANA_API_DASHBOARD_SAVE   = "/api/dashboards/db"
	GRAFANA_API_DASHBOARD_GET    = "/api/dashboards/uid/%s"
	GRAFANA_API_DASHBOARD_LIST   = "/api/search"
	GRAFANA_API_ANNOTAION_LIST   = "/api/annotations"
	GRAFANA_API_ALERTS_TEST      = "/api/alerts/test"
	GRAFANA_API_DATASOURCE_QUERY = "/api/datasources/proxy/1/api/v1/query_range"
)
