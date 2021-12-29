package events

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var testEvents = Events[string]{
	"user_created": []Listener[string]{
		func(payload string) {
			fmt.Printf("A new User just created!\n")
		},
		func(payload string) {
			fmt.Printf("A new User just created, *from second event listener\n")
		},
	},
	"user_joined": []Listener[string]{func(payload string) {
		if user, room, ok := strings.Cut(payload, " "); ok {
			fmt.Printf("%s joined to room: %s\n", user, room)
		}
	}},
	"user_left": []Listener[string]{func(payload string) {
		if user, room, ok := strings.Cut(payload, " "); ok {
			fmt.Printf("%s left from the room: %s\n", user, room)
		}
	}},
}

func createUser(evts EventEmitter[string], user string) {
	evts.Emit("user_created", user)
}

func joinUserTo(evts EventEmitter[string], user string, room string) {
	evts.Emit("user_joined", user+" "+room)
}

func leaveFromRoom(evts EventEmitter[string], user string, room string) {
	evts.Emit("user_left", user+" "+room)
}

func ExampleEvents() {
	evts := New[string]()

	// regiter our events to the default event emmiter
	for evt, listeners := range testEvents {
		evts.On(evt, listeners...)
	}

	user := "user1"
	room := "room1"

	createUser(evts, user)
	joinUserTo(evts, user, room)
	leaveFromRoom(evts, user, room)

	// Output:
	// A new User just created!
	// A new User just created, *from second event listener
	// user1 joined to room: room1
	// user1 left from the room: room1
}

func TestEvents(t *testing.T) {
	e := New[string]()
	expectedPayload := "this is my payload"

	e.On("my_event", func(payload string) {
		if payload != expectedPayload {
			t.Fatalf("Eexpected %s, got: %s", expectedPayload, payload)
		}
	})

	e.Emit("my_event", expectedPayload)
	if e.Len() != 1 {
		t.Fatalf("Length of the events is: %d, while expecting: %d", e.Len(), 1)
	}

	if e.Len() != 1 {
		t.Fatalf("Length of the listeners is: %d, while expecting: %d", e.ListenerCount("my_event"), 1)
	}

	e.RemoveAllListeners("my_event")
	if e.Len() != 0 {
		t.Fatalf("Length of the events is: %d, while expecting: %d", e.Len(), 0)
	}

	if e.Len() != 0 {
		t.Fatalf("Length of the listeners is: %d, while expecting: %d", e.ListenerCount("my_event"), 0)
	}
}

func TestEventsOnce(t *testing.T) {
	// on default
	Clear()

	var count = 0
	Once("my_event", func(payload interface{}) {
		if count > 0 {
			t.Fatalf("Once's listener fired more than one time! count: %d", count)
		}
		count++
	})

	if l := ListenerCount("my_event"); l != 1 {
		t.Fatalf("Real  event's listeners should be: %d but has: %d", 1, l)
	}

	if l := len(Listeners("my_event")); l != 1 {
		t.Fatalf("Real  event's listeners (from Listeners) should be: %d but has: %d", 1, l)
	}

	for i := 0; i < 10; i++ {
		Emit("my_event", nil)
	}

	time.Sleep(10 * time.Millisecond)

	if l := ListenerCount("my_event"); l > 0 {
		t.Fatalf("Real event's listeners length count should be: %d but has: %d", 0, l)
	}

	if l := len(Listeners("my_event")); l > 0 {
		t.Fatalf("Real event's listeners length count ( from Listeners) should be: %d but has: %d", 0, l)
	}

}

func TestRemoveListener(t *testing.T) {
	// on default
	e := New[string]()

	var count = 0
	listener := func(payload string) {
		if count > 1 {
			t.Fatal("Event listener should be removed")
		}

		count++
	}

	e.AddListener("my_event", listener)
	e.AddListener("my_event", func(payload string) {})
	e.AddListener("another_event", func(payload string) {})

	e.Emit("my_event", "test")

	if e.RemoveListener("my_event", listener) != true {
		t.Fatal("Should return 'true' when removes found listener")
	}

	if e.RemoveListener("foo_bar", listener) != false {
		t.Fatal("Should return 'false' when removes nothing")
	}

	if e.Len() != 2 {
		t.Fatal("Length of all listeners must be 2")
	}

	if e.ListenerCount("my_event") != 1 {
		t.Fatal("Length of 'my_event' event listeners must be 1")
	}

	e.Emit("my_event", "test")
}
