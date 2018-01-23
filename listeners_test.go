package cgoemitter_test

import (
	"testing"

	"github.com/supermock/cgoemitter"
)

var listenersList = cgoemitter.Listeners{
	cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
	cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
	cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
	cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
}

func TestAddListener(t *testing.T) {
	listeners := make(cgoemitter.Listeners, 0)

	for _, listener := range listenersList {
		listeners.AddListener(listener)
	}

	if len(listeners) != len(listenersList) {
		t.Errorf("Listener list was incorrect, got: %d, want: %d.", len(listeners), len(listenersList))
	}
}

func TestRemoveListener(t *testing.T) {
	listeners := make(cgoemitter.Listeners, 0)

	for _, listener := range listenersList {
		listeners.AddListener(listener)
		listeners.RemoveListener(listener)
	}

	if len(listeners) != 0 {
		t.Errorf("Listener list was incorrect, got: %d, want: %d.", len(listeners), 0)
	}
}
