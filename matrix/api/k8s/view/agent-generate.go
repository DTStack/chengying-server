package view

type AgentGenerateReq struct {
	Namespace string	`json:"namespace"`
	RegistryId int	`json:"registry_id"`
}

type AgentGenerateRsp struct {
	Yaml   string	`json:"yaml"`
}
