package strategy

const (
	STR_TYPE_HOST     = 0
	STR_TYPE_SERVICE  = 1
	CRON_TYPE_MINUTES = 0
	CRON_TYPE_HOURS   = 1
	CRON_TYPE_DAYS    = 2
)

const (
	MAX_HOST_TASK_NUM    = 10
	MAX_SERVICE_TASK_NUM = 1
)

type TaskInfo struct {
	TaskType          int
	StrategyId        int
	StrategyName      string
	SidecarId         string
	AgentId           string
	ExecScript        string
	Timeout           string
	Parameter         string
	Host              string
	ParentProductName string
	ProductName       string
	ServiceName       string
}
