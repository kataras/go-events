[Travis Widget]: https://img.shields.io/travis/kataras/go-events.svg?style=flat-square
[Travis]: http://travis-ci.org/kataras/go-events
[License Widget]: https://img.shields.io/badge/license-MIT%20%20License%20-E91E63.svg?style=flat-square
[License]: https://github.com/kataras/go-events/blob/master/LICENSE
[Release Widget]: https://img.shields.io/badge/release-v0.0.1-blue.svg?style=flat-square
[Release]: https://github.com/kataras/go-events/releases
[Chat Widget]: https://img.shields.io/badge/community-chat-00BCD4.svg?style=flat-square
[Chat]: https://kataras.rocket.chat/channel/go-events
[ChatMain]: https://kataras.rocket.chat/channel/go-events
[ChatAlternative]: https://gitter.im/kataras/go-events
[Report Widget]: https://img.shields.io/badge/report%20card-A%2B-F44336.svg?style=flat-square
[Report]: http://goreportcard.com/report/kataras/go-events
[Documentation Widget]: https://img.shields.io/badge/docs-reference-5272B4.svg?style=flat-square
[Documentation]: https://godoc.org/github.com/kataras/go-events
[Language Widget]: https://img.shields.io/badge/powered_by-Go-3362c2.svg?style=flat-square
[Language]: http://golang.org
[Platform Widget]: https://img.shields.io/badge/platform-All-yellow.svg?style=flat-square
[Awesome Widget]: https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg?style=flat-square
[Awesome WidgetAlternative]: https://img.shields.io/badge/awesome-%E2%9C%93-ff69b4.svg?style=flat-square
[Awesome]: https://github.com/avelino/awesome-go

<p align="center">
  <img src="/logo.jpg" height="400">
</p>

# Events [![Travis Widget]][Travis] [![Awesome Widget]][Awesome] [![License Widget]][License] [![Release Widget]][Release]

Simple nodejs-style EventEmmiter for Go Programming Language.



Quick view
------------

- `New` returns a new, empty EventEmmiter
- `On` is the func which registers the event listeners for a specific event
- `Emit` fires a particular event, this will call all functions(listeners) registered to this particular event
- `Remove` remove all registered listeners from a particular event
- `Len` returns the length of all registered events
- `LenListeners` returns the length of all registered listeners to a particular event

```go
import "github.com/kataras/go-events"

// initialize a new EventEmmiter to use
e := events.New()

// register an event with name "my_event" and one listener
e.On("my_event", func(payload ...interface{}) {
  message := payload[0].(string)
  print(message) // prints "this is my payload"
})

// fire the 'my_event' event
e.Emit("my_event", "this is my payload")

```

Default/global EventEmmiter
```go

// register an event with name "my_event" and one listener to the global(package level) default EventEmmiter
events.On("my_event", func(payload ...interface{}) {
  message := payload[0].(string)
  print(message) // prints "this is my payload"
})

// fire the 'my_event' event
events.Emit("my_event", "this is my payload")

```

Remove an event

```go

events.On("my_event", func(payload ...interface{}) {
  // first listener...
},func (payload ...interface{}){
  // second listener...
})

println(events.Len()) // prints 1
println(events.LenListeners("my_event")) // prints 2

// Remove our event, when/if we don't need this or we want to clear all of its listeners
events.Remove("my_event")

println(events.Len()) // prints 0
println(events.LenListeners("my_event")) // prints 0


```
[![Documentation Widget]][Documentation] [![Chat Widget]][Chat]

Installation
------------

The only requirement is the [Go Programming Language](https://golang.org/dl).

```bash
$ go get -u github.com/kataras/go-events
```

[![Language Widget]][Language] ![Platform Widget]

FAQ
------------

Explore [these questions](https://github.com/kataras/go-events/issues?go-events=label%3Aquestion) or navigate to the [community chat][Chat].

Versioning
------------

Current: v0.0.1

Read more about Semantic Versioning 2.0.0

 - http://semver.org/
 - https://en.wikipedia.org/wiki/Software_versioning
 - https://wiki.debian.org/UpstreamGuide#Releases_and_Versions

People
------------

The author of go-events is [@kataras](https://github.com/kataras).

If you're **willing to donate**, feel free to send **any** amount through paypal

[![](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=makis%40ideopod%2ecom&lc=GR&item_name=Iris%20web%20framework&item_number=iriswebframeworkdonationid2016&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


Contributing
------------

If you are interested in contributing to the go-events project, please make a PR.

[![Report Widget]][Report]


README template
------------

https://github.com/kataras/github-go-readme


License
------------

This project is licensed under the MIT License.

License can be found [here](LICENSE).
