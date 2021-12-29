<p align="center">
  <img src="./logo.jpg" height="400">
  <br/>
</p>

[![build status](https://img.shields.io/github/workflow/status/kataras/go-events/CI/master?style=for-the-badge)](https://github.com/kataras/go-events/actions) [![chat](https://img.shields.io/gitter/room/events/community.svg?color=cc2b5e&logo=gitter&style=for-the-badge)](https://gitter.im/events/community) [![donate](https://img.shields.io/badge/support-Go--Events-blue.svg?style=for-the-badge&logo=paypal)](https://iris-go.com/donate)
<br/>

Simple EventEmitter for Go Programming Language 1.18+ (**Generics** support). Inspired by <a href="https://nodejs.org/api/events.html">Nodejs EventEmitter</a>.

Overview
------------
`New[T any]() EventEmitter[T]  // New returns a new, empty, EventEmitter of T`


```go
// AddListener is an alias for .On(eventName, listener).
AddListener(EventName, ...Listener)
// Emit fires a particular event,
// Synchronously calls each of the listeners registered for the event named
// eventName, in the order they were registered,
// passing the supplied arguments to each.
Emit(EventName, ...interface{})
// EventNames returns an array listing the events for which the emitter has registered listeners.
// The values in the array will be strings.
EventNames() []EventName
// GetMaxListeners returns the max listeners for this emmiter
// see SetMaxListeners
GetMaxListeners() int
// ListenerCount returns the length of all registered listeners to a particular event
ListenerCount(EventName) int
// Listeners returns a copy of the array of listeners for the event named eventName.
Listeners(EventName) []Listener
// On registers a particular listener for an event, func receiver parameter(s) is/are optional
On(EventName, ...Listener)
// Once adds a one time listener function for the event named eventName.
// The next time eventName is triggered, this listener is removed and then invoked.
Once(EventName, ...Listener)
// RemoveAllListeners removes all listeners, or those of the specified eventName.
// Note that it will remove the event itself.
// Returns an indicator if event and listeners were found before the remove.
RemoveAllListeners(EventName) bool
// Clear removes all events and all listeners, restores Events to an empty value
Clear()
// SetMaxListeners obviously this function allows the MaxListeners
// to be decrease or increase. Set to zero for unlimited
SetMaxListeners(int)
// Len returns the length of all registered events
Len() int
```


```go
import "github.com/kataras/go-events"

// initialize a new EventEmitter to use
e := events.New[string]()

// register an event with name "my_event" and one listener
e.On("my_event", func(message string) {
  print(message) // prints "this is my payload"
})

// fire the 'my_event' event
e.Emit("my_event", "some message")

```

Default/global EventEmitter
```go

// register an event with name "my_event" and one listener to the global(package level) default EventEmitter
events.On("my_event", func(payload interface{}) {
  message := payload.(string)
  print(message) // prints "this is my payload"
})

// fire the 'my_event' event
events.Emit("my_event", "this is my payload")

```

Remove an event

```go
e := New[int]()

e.On("my_event", func(payload int) {
  // first listener...
},func (payload int){
  // second listener...
})

println(e.Len()) // prints 1
println(e.ListenerCount("my_event")) // prints 2

// Remove our event, when/if we don't need this or we want to clear all of its listeners
e.RemoveAllListeners("my_event")

println(e.Len()) // prints 0
println(e.ListenerCount("my_event")) // prints 0


```
Installation
------------

The only requirement is the [Go Programming Language](https://golang.org/dl), at least version 1.18.

```bash
$ go get github.com/kataras/go-events@latest
```


FAQ
------------

Explore [these questions](https://github.com/kataras/go-events/issues?go-events=label%3Aquestion) or navigate to the [community chat][Chat].

Versioning
------------

Current: v1.0.0

Read more about Semantic Versioning 2.0.0

 - http://semver.org/
 - https://en.wikipedia.org/wiki/Software_versioning
 - https://wiki.debian.org/UpstreamGuide#Releases_and_Versions

People
------------

The author of go-events is [@kataras](https://github.com/kataras).

Contributing
------------

If you are interested in contributing to the go-events project, please make a PR.

License
------------

This project is licensed under the MIT License.

License can be found [here](LICENSE).
