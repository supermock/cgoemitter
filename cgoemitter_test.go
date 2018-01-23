package cgoemitter_test

import (
	"testing"

	"github.com/supermock/cgoemitter"
)

func TestOn(t *testing.T) {
	for _, event := range eventsList {
		for _, listener := range event.listeners {
			cgoemitter.On(event.eventName, listener)
		}
	}

	for _, event := range eventsList {
		if listeners, err := cgoemitter.GetListeners(event.eventName); err != nil {
			t.Errorf("cgoemitter.GetListeners() was incorrect, got: %v, want: %v.", err, nil)
		} else {
			if len(listeners) != len(event.listeners) {
				t.Errorf("cgoemitter.GetListeners() was incorrect, got: %d, want: %d.", len(listeners), len(event.listeners))
			}
		}
	}
}

func TestOff(t *testing.T) {
	for _, event := range eventsList {
		for _, listener := range event.listeners {
			cgoemitter.Off(event.eventName, listener)
		}

		if listeners, err := cgoemitter.GetListeners(event.eventName); err != nil {
			t.Errorf("cgoemitter.GetListeners() was incorrect, got: %v, want: %v.", err, nil)
		} else {
			if len(listeners) != 0 {
				t.Errorf("cgoemitter.GetListeners() was incorrect, got: %d, want: %d.", len(listeners), 0)
			}
		}
	}
}

func TestGetListeners(t *testing.T) {
	if _, err := cgoemitter.GetListeners("null"); err == nil {
		t.Errorf("cgoemitter.GetListeners() was incorrect, got: %t, want: %t.", err == nil, false)
	}
}
