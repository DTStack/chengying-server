package event

const (
	maxEventCache = 128
)

var eventManager *eventManage

func init() {
	eventManager = &eventManage{eventQueue: make(chan *Event, maxEventCache)}
	go eventManager.EventHandler()
}

type eventManage struct {
	Observable
	eventQueue chan *Event
}

type EventManager interface {
	EventReciever(event *Event)
	EventHandler()
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	Notify(event *Event)
}

func (this *eventManage) AddObserver(observer Observer) {
	this.RegistObserver(observer)
}

func (this *eventManage) RemoveObserver(observer Observer) {
	this.UnregistObserver(observer)
}

func (this *eventManage) Notify(event *Event) {
	this.NofityObservers(*event)
}

func (this *eventManage) EventHandler() {
	for {
		ev, ok := <-this.eventQueue

		if !ok {
			return
		}
		this.Notify(ev)
	}
}

func (this *eventManage) EventReciever(event *Event) {
	// log.Debugf("[EventReciever] receive: %+v", event.Data)
	this.eventQueue <- event
}

func GetEventManager() EventManager {
	return eventManager
}
