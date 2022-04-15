package event

// Observer pattern
// https://en.wikipedia.org/wiki/Observer_pattern
type Observer interface {
	Notify(event Event)
}

type Observable struct {
	Observers []Observer
}

func (self *Observable) RegistObserver(observer Observer) {
	self.Observers = append(self.Observers, observer)
}

func (self *Observable) UnregistObserver(observer Observer) {
	for index, ob := range self.Observers {
		if ob == observer {
			self.Observers = append(self.Observers[:index], self.Observers[index+1:]...)
			break
		}
	}
}

func (self Observable) NofityObservers(event Event) {
	for observer := range self.Observers {
		self.Observers[observer].Notify(event)
	}
}
