package event

type Event struct {
	AgentId string
	Type    string
	Data    interface{}
}

//event report from agent server

const (
	REPORT_EVENT_HEALTH_CHECK         = "report.event.instance.healthcheck"
	REPORT_EVENT_HEALTH_CHECK_CANCEL  = "report.event.instance.healthcheck.cancel"
	REPORT_EVENT_INSTANCE_ERROR       = "report.event.instance.error"
	REPORT_EVENT_INSTANCE_PERFORMANCE = "report.event.instance.error"
)
