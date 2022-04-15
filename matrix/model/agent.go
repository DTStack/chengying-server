package model

import (
	"strings"
	"time"
)

type Agent struct {
	Id             string     `db:"id"`
	SidecarId      string     `db:"sidecar_id"`
	Type           int        `db:"type"`
	Name           string     `db:"name"`
	Version        string     `db:"version"`
	IsUninstalled  int        `db:"is_uninstalled"`
	DeployDate     *time.Time `db:"deploy_date"`
	AutoDeployment int        `db:"auto_deployment"`
	LastUpdateDate *time.Time `db:"last_update_date"`
	AutoUpdated    int        `db:"auto_updated"`
}

type ExecScriptResponse struct {
	Seqno    uint32 `json:"seqno,omitempty"`
	Failed   bool   `json:"failed,omitempty"`
	Response string `json:"response,omitempty"`
	AgentId  string `json:"agentId"`
}

func GetAgentByServices(names []string) ([]Agent, error) {
	for i := range names {
		names[i] = "\"" + names[i] + "\""
	}
	query := "SELECT * FROM agent_list WHERE name IN (" + strings.Join(names, ",") + ")"
	agents := make([]Agent, 0)

	err := USE_MYSQL_DB().Select(&agents, query)
	return agents, err
}
