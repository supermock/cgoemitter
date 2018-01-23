package cgoemitter

import "sync"

//Events | List of events
type Events map[string]*Listeners

var eventsMutex = &sync.Mutex{}

//AddEvent | Add a new event
func (events *Events) AddEvent(eventName string) *Listeners {
	listeners := make(Listeners, 0)
	eventsMutex.Lock()
	(*events)[eventName] = &listeners
	eventsMutex.Unlock()
	return &listeners
}

//RemoveEvent | Remove an event
func (events *Events) RemoveEvent(eventName string) {
	eventsMutex.Lock()
	delete(*events, eventName)
	eventsMutex.Unlock()
}

//Has | Returns the value and if the event already exists it returns true
func (events Events) Has(eventName string) (*Listeners, bool) {
	eventsMutex.Lock()
	listeners, exist := events[eventName]
	eventsMutex.Unlock()
	return listeners, exist
}
