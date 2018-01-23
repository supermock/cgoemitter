package cgoemitter_test

import (
	"testing"

	"github.com/supermock/cgoemitter"
)

var eventsList = []struct {
	eventName string
	listeners cgoemitter.Listeners
}{
	{
		"event-1",
		cgoemitter.Listeners{
			cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
			cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
		},
	},
	{
		"event-2",
		cgoemitter.Listeners{
			cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
			cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
			cgoemitter.NewListener(func(args cgoemitter.Arguments) {}),
		},
	},
}

func TestAddEvent(t *testing.T) {
	events := make(cgoemitter.Events, 0)

	for _, event := range eventsList {
		events.AddEvent(event.eventName)
	}

	if len(events) != len(eventsList) {
		t.Errorf("Event list was incorrect, got: %d, want: %d.", len(events), len(eventsList))
	}
}

func TestRemoveEvent(t *testing.T) {
	events := make(cgoemitter.Events, 0)

	for _, event := range eventsList {
		events.AddEvent(event.eventName)
		events.RemoveEvent(event.eventName)
	}

	if len(events) != 0 {
		t.Errorf("Event list was incorrect, got: %d, want: %d.", len(events), 0)
	}
}

func TestHas(t *testing.T) {
	events := make(cgoemitter.Events, 0)

	for _, event := range eventsList {
		events.AddEvent(event.eventName)

		listeners, exists := events.Has(event.eventName)
		if !exists {
			t.Errorf("Event list was incorrect, got: %t, want: %t.", exists, true)
		}

		for _, listener := range event.listeners {
			listeners.AddListener(listener)
		}

		if len(*listeners) != len(event.listeners) {
			t.Errorf("Event listener list was incorrect, got: %d, want: %d.", len(*listeners), len(event.listeners))
		}

		for _, listener := range event.listeners {
			listeners.RemoveListener(listener)
		}

		if len(*listeners) != 0 {
			t.Errorf("Event listener list was incorrect, got: %d, want: %d.", len(*listeners), 0)
		}
	}
}
