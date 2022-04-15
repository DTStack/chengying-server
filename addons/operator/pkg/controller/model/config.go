package model

import (
	"fmt"
	"os"
)

const (
	CommonLogPath = "/var/log/commonlog"
	LogVolume     = "sidecarvolume"
	Configmap     = "log-configmap"
	LogConfigPath = "/etc/log_config"
	PromtailPort  = 3101
	LogSubpath    = "commonlog"
)

var (
	ENV_LOG_IMAGE          = "LOG_IMAGE"
	ENV_LOG_TYPE           = "LOG_TYPE"
	ENV_LOG_SERVER_ADDRESS = "LOG_SERVER_ADDRESS"
	ENV_LOG_MEM_LIMIT      = "LOG_MEM_LIMIT"
	ENV_LOG_CPU_LIMIT      = "LOG_CPU_LIMIT"
	ENV_LOG_MEM_REQUEST    = "LOG_MEM_REQUEST"
	ENV_LOG_CPU_REQUEST    = "LOG_CPU_REQUEST"
	ENV_LOG_SWITCH         = "LOG_SWITCH"
	LogArgs                []string
	LogImage               = os.Getenv(ENV_LOG_IMAGE)
	LogType                = os.Getenv(ENV_LOG_TYPE)
	LogServerAddress       = os.Getenv(ENV_LOG_SERVER_ADDRESS)
)

func init() {

	switch LogType {
	case "promtail":
		LogArgs = []string{
			fmt.Sprintf("-config.file=%s/promtail.yaml", LogConfigPath),
			fmt.Sprintf("-client.url=http://%s/loki/api/v1/push", LogServerAddress),
			"-config.expand-env=true"}
	case "filebeat":
		LogArgs = []string{
			"-c",
			fmt.Sprintf("%s/filebeat.yml", LogConfigPath),
			"-e"}
	}

}
