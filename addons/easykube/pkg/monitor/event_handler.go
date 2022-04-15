package monitor

import monitorevents "dtstack.com/dtstack/easymatrix/addons/easymonitor/pkg/monitor/events"

func GetEvents() []monitorevents.Eventer {
	events := EventCache.Get()
	return events
}
