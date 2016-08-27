package events

type (
	// EventListeners the listeners type, it's just a []func(...interface{})
	EventListeners []func(...interface{})
	// Events the type for registered listeners, it's just a map[string][]func(...interface{})
	Events map[string]EventListeners

	// EventEmmiter is the message/or/event manager
	EventEmmiter interface {
		// On registers a particular listener for an event, func receiver parameter(s) is/are optional
		On(string, ...func(...interface{}))
		// Emit fires a particular event, second parameter(s) is/are optional, if filled then the event has some information to send to the listener
		Emit(string, ...interface{})
		// Remove receives an event and destroys/ de-activates/unregisters the listeners belongs to it( the event)  and removes itself from the events
		// returns true if something has been removed, otherwise false
		Remove(string) bool
		// Len returns the length of all registered events
		Len() int
		// LenListeners returns the length of all registered listeners to a particular event
		LenListeners(string) int
	}

	emmiter struct {
		evtListeners Events
	}
)

func (e Events) copyTo(emmiter EventEmmiter) {
	if e != nil && len(e) > 0 {
		// register the events to/with their listeners
		for evt, listeners := range e {
			if len(listeners) > 0 {
				for i := range listeners {
					emmiter.On(evt, listeners[i])
				}
			}
		}
	}
}

// New returns a new, empty EventEmmiter
func New() EventEmmiter {
	return &emmiter{}
}

var (
	_              EventEmmiter = &emmiter{}
	defaultEmmiter              = New()
)

// On registers a particular listener for an event, func receiver parameter(s) is/are optional
func On(evt string, listener ...func(data ...interface{})) {
	defaultEmmiter.On(evt, listener...)
}

func (e *emmiter) On(evt string, listener ...func(data ...interface{})) {
	if e.evtListeners == nil {
		e.evtListeners = Events{}
	}
	if e.evtListeners[evt] == nil {
		e.evtListeners[evt] = EventListeners{}
	}
	e.evtListeners[evt] = append(e.evtListeners[evt], listener...)
}

// Emit fires a particular event, second parameter(s) is/are optional, if filled then the event has some information to send to the listener
func Emit(evt string, data ...interface{}) {
	defaultEmmiter.Emit(evt, data...)
}

func (e *emmiter) Emit(evt string, data ...interface{}) {
	if e.evtListeners == nil {
		return // has no listeners to emit/speak yet
	}
	if listeners := e.evtListeners[evt]; listeners != nil && len(listeners) > 0 { // len() should be just fine, but for any case on future...
		for i := range listeners {
			l := listeners[i]
			l(data...)
		}
	}
}

// Remove receives an event and destroys/ de-activates/unregisters the listeners belongs to it( the event) and removes itself from the events
// returns true if something has been removed, otherwise false
func Remove(evt string) bool {
	return defaultEmmiter.Remove(evt)
}

func (e *emmiter) Remove(evt string) bool {
	if e.evtListeners == nil {
		return false // has nothing to remove
	}

	if listeners := e.evtListeners[evt]; listeners != nil {
		//e.evtListeners[evt] = EventListeners{}
		delete(e.evtListeners, evt)
		if len(listeners) > 0 {
			return true
		}
	}

	return false
}

// Len returns the length of all registered events
func Len() int {
	return defaultEmmiter.Len()
}

// LenListeners returns the length of all registered listeners to a particular event
func LenListeners(evt string) int {
	return defaultEmmiter.LenListeners(evt)
}

func (e *emmiter) Len() int {
	if e.evtListeners == nil {
		return 0
	}
	return len(e.evtListeners)
}

func (e *emmiter) LenListeners(evt string) int {
	if e.evtListeners == nil {
		return 0
	}

	if listeners := e.evtListeners[evt]; listeners != nil { // len() should be just fine, but for any case on future...
		return len(listeners)
	}

	return 0
}
