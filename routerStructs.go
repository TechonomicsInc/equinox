package equinox

import (
    "sync"
)

// Handler defines the basic layout of Handler structs
type Handler interface {
    // Called after the handler is registered through AddRoute()
    Init(
        session interface{},
    )

    // Called to retrieve the Listener patterns
    Listeners() []*Listener

    // Called if any of the listeners matched
    Action(
        command string,
        content string,
        params map[string]string,
        wmsg interface{},
        wsession interface{},
    )
}

// Listener is a simple container that may either hold a string-pattern or a RegExp object
type Listener struct {
    Content  string
}

type RegexListener struct {
    Listener
}

// TODO: write docs about the most important thing in equinox ._.
type Router struct {
    sync.RWMutex

    debugMode bool

    lastResort        POGOFuncW1
    prefixHandler     PrefixHandler
    panicHandler      PanicHandler
    parseErrorHandler ParseErrorHandler

    EventHandlers map[Event][]AdapterFunc
    Routes        map[*Listener][]Handler
}
