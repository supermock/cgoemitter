package cgoemitter

/*
#include <stdlib.h>
#include "cgoemitter.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

var events *Events

var (
	//ErrUnknownEvent | This event does not exist
	ErrUnknownEvent = errors.New("This event does not exist")
)

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

func dispatchEvent(listeners *Listeners, args Arguments) {
	for _, listener := range *listeners {
		(*listener)(args)
	}
}

//export emit
func emit(eventName *C.char, cgoEmitterArgs *C.struct_cgoemitter_args_t) {
	listeners := loadListeners(C.GoString(eventName), false)

	var args Arguments
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&args))
	sliceHeader.Cap = int(cgoEmitterArgs.args_cap)
	sliceHeader.Len = int(cgoEmitterArgs.args_len)
	sliceHeader.Data = uintptr(*cgoEmitterArgs.args)
	defer C.free(unsafe.Pointer(cgoEmitterArgs.args))
	defer args.free()

	if listeners != nil {
		dispatchEvent(listeners, args)
	} else {
		listeners = loadListeners("cgoemitter-warnings", false)

		if listeners != nil {
			warningMessage := unsafe.Pointer(C.CString(fmt.Sprintf("The '%s' event was triggered by C, but there are no handlers", C.GoString(eventName))))
			defer C.free(warningMessage)

			warningArgs := make(Arguments, 0, 1)
			warningArgs = append(warningArgs, warningMessage)

			dispatchEvent(listeners, warningArgs)
		}
	}
}

//On | Adds a new event if it does not exist, and also adds a new listener
func On(eventName string, listener *ListenerFunc) {
	listeners := loadListeners(eventName, true)
	listeners.AddListener(listener)
}

//Off | Removes a listener from the event or the event itself if the listener number equals zero
func Off(eventName string, listener *ListenerFunc) {
	listeners := loadListeners(eventName, false)
	if listeners != nil {
		listeners.RemoveListener(listener)

		if len(*listeners) == 0 {
			events.RemoveEvent(eventName)
		}
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
	return Listeners{}, ErrUnknownEvent
}
