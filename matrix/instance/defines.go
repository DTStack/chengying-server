package instance

type NewInstancerParam struct {
	Pid         int    `json:"pid"`
	Ip          string `json:"ip"`
	ServiceName string `json:"name"`
	Schema      string `json:"schema"`
}

const (
	LINUX_EXEC_HEADER = "#!/bin/bash\n"
	EXEC_TIMEOUT      = "1h"
)
