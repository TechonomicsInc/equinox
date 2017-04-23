package equinox

import (
    "code.lukas.moe/x/wormhole"
    "sync"
)

// Handler defines the basic layout of Handler structs
type Handler interface {
    // Called after the handler is registered through AddRoute()
    Init()

    // Called to retrieve the Listener patterns
    Listeners() []*Listener

    // Called if any of the listeners matched
    Action(
        command string,
        content string,
        params map[string]string,
        msg *wormhole.Wormhole,
    )
}

// Listener is a simple container that may either hold a string-pattern or a RegExp object
type Listener struct {
    IsRegexp bool
    Content  string
}

type Router struct {
    sync.RWMutex

    debugMode bool

    lastResort        POGOFunc
    prefixHandler     PrefixHandler
    panicHandler      PanicHandler
    parseErrorHandler ParseErrorHandler

    EventHandlers map[Event][]AdapterFunc
    Routes        map[*Listener][]Handler
}
