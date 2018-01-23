package cgoemitter

import "sync"

//ListenerFunc | Listener function
type ListenerFunc func(Arguments)

//Listeners | List of listeners
type Listeners []*ListenerFunc

var listenersMutex = &sync.Mutex{}

//AddListener | Adds a new listener
func (listeners *Listeners) AddListener(listener *ListenerFunc) {
	listenersMutex.Lock()
	*listeners = append(*listeners, listener)
	listenersMutex.Unlock()
}

//RemoveListener | Remove a listener
func (listeners *Listeners) RemoveListener(listener *ListenerFunc) {
	listenersMutex.Lock()
	for i, listenerItem := range *listeners {
		if listenerItem == listener {
			*listeners = append((*listeners)[:i], (*listeners)[i+1:]...)
			break
		}
	}
	listenersMutex.Unlock()
}
