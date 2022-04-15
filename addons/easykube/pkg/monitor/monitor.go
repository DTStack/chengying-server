package monitor

import (
	easymonitor "dtstack.com/dtstack/easymatrix/addons/easymonitor/pkg"
	monitorevents "dtstack.com/dtstack/easymatrix/addons/easymonitor/pkg/monitor/events"
)

func StartMonitor(namespace string, ch chan struct{}) error {
	transmitor := &AgentModeTransmitor{}
	monitorevents.Transmitor = transmitor
	err := easymonitor.StartMonitorController("", "", namespace, ch)
	if err != nil {
		return err
	}
	return nil
}

type AgentModeTransmitor struct {
}

func (a *AgentModeTransmitor) Push(event monitorevents.Eventer) {
	e := event.(*monitorevents.Event)
	EventCache.Push(e)
}

func (a *AgentModeTransmitor) Process() {

}
