package monitor

import (
	monitorevents "dtstack.com/dtstack/easymatrix/addons/easymonitor/pkg/monitor/events"
	"sync"
)

var EventCache = &Cache{
	Events: make([]monitorevents.Eventer, 0, 1024),
}

type Cache struct {
	Events []monitorevents.Eventer
	mu     sync.Mutex
}

func (c *Cache) Push(evnet monitorevents.Eventer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Events = append(c.Events, evnet)
}

func (c *Cache) Get() []monitorevents.Eventer {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.Events) == 0 {
		return nil
	}
	events := c.Events
	c.Events = make([]monitorevents.Eventer, 0, 1024)
	return events
}
