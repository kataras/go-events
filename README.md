<a href="https://travis-ci.org/kataras/go-events"><img src="https://img.shields.io/travis/kataras/go-events.svg?style=flat-square" alt="Build Status"></a>
<a href="https://github.com/kataras/go-events/blob/master/LICENSE"><img src="https://img.shields.io/badge/%20license-MIT%20%20License%20-E91E63.svg?style=flat-square" alt="License"></a>
<a href="https://github.com/kataras/go-events/releases"><img src="https://img.shields.io/badge/%20release%20-%20v0.0.1-blue.svg?style=flat-square" alt="Releases"></a>
<a href="#docs"><img src="https://img.shields.io/badge/%20docs-reference-5272B4.svg?style=flat-square" alt="Read me docs"></a>
<a href="https://kataras.rocket.chat/channel/go-events"><img src="https://img.shields.io/badge/%20community-chat-00BCD4.svg?style=flat-square" alt="Build Status"></a>
<a href="https://golang.org"><img src="https://img.shields.io/badge/powered_by-Go-3362c2.svg?style=flat-square" alt="Built with GoLang"></a>
<a href="#"><img src="https://img.shields.io/badge/platform-Any--OS-yellow.svg?style=flat-square" alt="Platforms"></a>


Simple EventEmmiter for Go Programming Language.


Installation
------------
The only requirement is the [Go Programming Language](https://golang.org/dl).

```bash
$ go get -u github.com/kataras/go-events
```


Docs
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

FAQ
------------
Explore [these questions](https://github.com/kataras/go-events/issues?go-events=label%3Aquestion) or navigate to the [community chat][Chat].

Versioning
------------

Current: **v0.0.1**



People
------------
The author of go-events is [@kataras](https://github.com/kataras).

If you're **willing to donate**, feel free to send **any** amount through paypal

[![](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=makis%40ideopod%2ecom&lc=GR&item_name=Iris%20web%20framework&item_number=iriswebframeworkdonationid2016&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


Contributing
------------
If you are interested in contributing to the go-events project, please make a PR.

License
------------

This project is licensed under the MIT License.

License can be found [here](LICENSE).

[Travis Widget]: https://img.shields.io/travis/kataras/go-events.svg?style=flat-square
[Travis]: http://travis-ci.org/kataras/go-events
[License Widget]: https://img.shields.io/badge/license-MIT%20%20License%20-E91E63.svg?style=flat-square
[License]: https://github.com/kataras/go-events/blob/master/LICENSE
[Release Widget]: https://img.shields.io/badge/release-v4.1.1-blue.svg?style=flat-square
[Release]: https://github.com/kataras/go-events/releases
[Chat Widget]: https://img.shields.io/badge/community-chat-00BCD4.svg?style=flat-square
[Chat]: https://kataras.rocket.chat/channel/go-events
[ChatMain]: https://kataras.rocket.chat/channel/go-events
[ChatAlternative]: https://gitter.im/kataras/go-events
[Report Widget]: https://img.shields.io/badge/report%20card-A%2B-F44336.svg?style=flat-square
[Report]: http://goreportcard.com/report/kataras/go-events
[Documentation Widget]: https://img.shields.io/badge/documentation-reference-5272B4.svg?style=flat-square
[Documentation]: https://www.gitbook.com/book/kataras/go-events/details
[Language Widget]: https://img.shields.io/badge/powered_by-Go-3362c2.svg?style=flat-square
[Language]: http://golang.org
[Platform Widget]: https://img.shields.io/badge/platform-Any--OS-gray.svg?style=flat-square
