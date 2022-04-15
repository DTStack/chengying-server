package host

import (
	"testing"
)

func TestAgentInstall_GetAgentInstallCmd(t *testing.T) {
	err, cmd := AgentInstall.GetAgentInstallCmd(2, "172.16.8.175:8864", "", "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("agent install cmd: %#v", cmd)
}
