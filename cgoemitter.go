package cgoemitter

/*
#include <stdlib.h>
#include "cgoemitter.h"
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

var events *Events

func init() {
	eventsList := make(Events, 0)
	events = &eventsList
}

func loadListeners(eventName string, create bool) *Listeners {
	listeners, exist := events.Has(eventName)
	if !exist && create {
		return events.AddEvent(eventName)
	}
	return listeners
}

//export emit
func emit(eventName *C.char, cgoEmitterArgs *C.struct_cgoemitter_args_t) {
	listeners := loadListeners(C.GoString(eventName), true)

	var args Arguments
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&args))
	sliceHeader.Cap = int(cgoEmitterArgs.args_cap)
	sliceHeader.Len = int(cgoEmitterArgs.args_len)
	sliceHeader.Data = uintptr(*cgoEmitterArgs.args)
	defer C.free(unsafe.Pointer(cgoEmitterArgs.args))
	defer args.free()

	for _, listener := range *listeners {
		(*listener)(args)
	}
}

//On | Adds a new event if it does not exist, and also adds a new listener
func On(eventName string, listener *ListenerFunc) {
	listeners := loadListeners(eventName, true)
	listeners.AddListener(listener)
}

//Off | Removes a listener from the event
func Off(eventName string, listener *ListenerFunc) {
	listeners := loadListeners(eventName, false)
	if listeners != nil {
		listeners.RemoveListener(listener)
	}
}

//NewListener | Instance a new listener
func NewListener(listener ListenerFunc) *ListenerFunc {
	return &listener
}

//GetListeners | Return all listeners for an event
func GetListeners(eventName string) (Listeners, error) {
	listeners := loadListeners(eventName, false)
	if listeners != nil {
		return *listeners, nil
	}
	return Listeners{}, errors.New("This event does not exist")
}
