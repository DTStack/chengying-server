package agent

const (
	AGENT_INSTALL_URI           = "/api/v1/agent/install"
	AGENT_INSTALL_SYNC_URI      = "/api/v1/agent/installSync"
	AGENT_UNINSTALL_URI         = "/api/v1/agent/%s/uninstall"
	AGENT_UNINSTALL_SYNC_URI    = "/api/v1/agent/%s/uninstallSync"
	AGENT_CANCEL_SYNC_URI       = "/api/v1/agent/cancelOperation"
	AGENT_START_URI             = "/api/v1/agent/%s/start"
	AGENT_START_SYNC_URI        = "/api/v1/agent/%s/startSync"
	AGENT_START_SYNC_PARAMS_URI = "/api/v1/agent/%s/startSyncWithParam"
	AGENT_STOP_URI              = "/api/v1/agent/%s/stop"
	AGENT_STOP_SYNC_URI         = "/api/v1/agent/%s/stopSync"
	AGENT_EXEC_SYNC_URI         = "/api/v1/sidecar/%s/execscriptSync"
	AGENT_REST_SYNC_URI         = "/api/v1/sidecar/%s/execrestSync"
	AGENT_EXEC_OFTEN_SYNC_URI   = "/api/v1/sidecar/%s/execscriptOftenSync"
	AGENT_CONFIG_SYNC_URI       = "/api/v1/agent/%s/configSync"
)
