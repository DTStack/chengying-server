package agent

type InstanceEvent struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (this *InstanceEvent) GetType() string {
	return this.Type
}

func (this *InstanceEvent) GetMessage() string {
	return this.Message
}

type InstallEvent struct {
	InstanceEvent
	InstallSchema  interface{}   `json:"installSchema"`
	InstallParam   *InstallParms `json:"installParam"`
	InstallResp    interface{}   `json:"installResp"`
	ConfigParam    []interface{} `json:"configParam"`
	ConfigResp     []interface{} `json:"configResp"`
	PostDeployResp interface{}   `json:"configUpdateResp"`
}

type UnInstallEvent struct {
	InstanceEvent
	UnInstallParam *ShellParams `json:"unInstallParam"`
	UnInstallResp  interface{}  `json:"unInstallResp"`
}

type ConfigEvent struct {
	InstanceEvent
	ConfigSchema interface{}   `json:"configSchema"`
	ConfigPath   []string      `json:"configPath"`
	ConfigResp   []interface{} `json:"configResp"`
}

type StartEvent struct {
	InstanceEvent
	StartParam *StartParams `json:"startParam"`
	StartResp  interface{}  `json:"startResp"`
}

type StopEvent struct {
	InstanceEvent
	StopResp interface{} `json:"startResp"`
}

type ExecEvent struct {
	InstanceEvent
	ExecScriptParam *ExecScriptParams `json:"execScriptParam"`
	ExecResp        interface{}       `json:"execResp"`
}

type ErrorEvent struct {
	InstanceEvent
	ErrorResp interface{} `json:"errorResp"`
}

type UnknownEvent struct {
	InstanceEvent
}
