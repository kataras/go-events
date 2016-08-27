package events

import (
	"fmt"
	"testing"
)

var testEvents = Events{
	"user_created": EventListeners{
		func(payload ...interface{}) {
			fmt.Printf("A new User just created!\n")
		},
		func(payload ...interface{}) {
			fmt.Printf("A new User just created, *from second event listener\n")
		},
	},
	"user_joined": EventListeners{func(payload ...interface{}) {
		user := payload[0].(string)
		room := payload[1].(string)
		fmt.Printf("%s joined to room: %s\n", user, room)
	}},
	"user_left": EventListeners{func(payload ...interface{}) {
		user := payload[0].(string)
		room := payload[1].(string)
		fmt.Printf("%s left from the room: %s\n", user, room)
	}},
}

func createUser(user string) {
	Emit("user_created", user)
}

func joinUserTo(user string, room string) {
	Emit("user_joined", user, room)
}

func leaveFromRoom(user string, room string) {
	Emit("user_left", user, room)
}

func ExampleEvents() {
	// regiter our events to the default event emmiter
	for evt, listeners := range testEvents {
		On(evt, listeners...)
	}

	user := "user1"
	room := "room1"

	createUser(user)
	joinUserTo(user, room)
	leaveFromRoom(user, room)

	// Output:
	// A new User just created!
	// A new User just created, *from second event listener
	// user1 joined to room: room1
	// user1 left from the room: room1
}

func TestEvents(t *testing.T) {
	e := New()
	expectedPayload := "this is my payload"

	e.On("my_event", func(payload ...interface{}) {
		if len(payload) <= 0 {
			t.Fatal("Expected payload but got nothing")
		}

		if s, ok := payload[0].(string); !ok {
			t.Fatalf("Payload is not the correct type, got: %#v", payload[0])
		} else if s != expectedPayload {
			t.Fatalf("Eexpected %s, got: %s", expectedPayload, s)
		}
	})

	e.Emit("my_event", expectedPayload)
	if e.Len() != 1 {
		t.Fatalf("Length of the events is: %d, while expecting: %d", e.Len(), 1)
	}

	if e.Len() != 1 {
		t.Fatalf("Length of the listeners is: %d, while expecting: %d", e.LenListeners("my_event"), 1)
	}

	e.Remove("my_event")
	if e.Len() != 0 {
		t.Fatalf("Length of the events is: %d, while expecting: %d", e.Len(), 0)
	}

	if e.Len() != 0 {
		t.Fatalf("Length of the listeners is: %d, while expecting: %d", e.LenListeners("my_event"), 0)
	}
}
