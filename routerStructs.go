package equinox

import (
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
        wmsg interface{},
        wsession interface{},
    )
}

type Listener struct {
    Content  string
    IsRegexp bool
}

func NewListener(content string) *Listener {
    return &Listener{
        Content: content,
        IsRegexp: false,
    }
}

func NewRegexListener(content string) *Listener {
    return &Listener{
        Content: content,
        IsRegexp: true,
    }
}

func BulkRegister(fn func(string)(*Listener), contents []string) []*Listener {
    listeners := []*Listener{}

    for _, content := range contents {
        listeners = append(listeners, fn(content))
    }

    return listeners
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
