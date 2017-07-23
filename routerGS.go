package equinox

import (
    "fmt"
    "strings"
    "code.lukas.moe/x/equinox/caches"
)

// NewRouter constructs a router object with some default adapters
func NewRouter() *Router {
    r := &Router{}

    r.Lock()
    r.EventHandlers = make(map[Event][]AdapterFunc)
    r.Routes = make(map[string]Handler)
    r.RouteMeta = make(map[string]ListenerMeta)
    r.Unlock()

    r.UseDebugMode(false)
    r.SetPanicHandler(DefaultPanicHandler)
    r.SetParseErrorHandler(DefaultParseErrorHandler)
    r.SetLastResort(NOOPW1)

    return r
}

// AddRoutes is syntactic sugar for a loop that calls AddRoute() multiple times.
func (r *Router) AddRoutes(handlers []Handler) {
    for _, h := range handlers {
        r.AddRoute(h)
    }
}

// AddRoute calls the handler's Init() and addss it's listeners to the routing table
func (r *Router) AddRoute(handler Handler) {
    r.Lock()
    defer r.Unlock()

    OnDebug(func() {
        logf("Registering handler %s", TypeOf(handler))
    })

    for _, l := range handler.Listeners() {
        parts := strings.Fields(l)
        l = parts[0]

        if len(parts) > 1 {
            parts = parts[1:]
        } else {
            parts = []string{}
        }

        OnDebug(func() {
            logf("--- Found listener: %s", l)
        })

        if _, ok := r.Routes[l]; !ok {
            r.Routes[l] = handler
            r.RouteMeta[l] = ListenerMeta{Expression: parts}
        } else {
            panic(fmt.Errorf(
                "Tried to add duplicate route %s with handler \n%#v",
                l,
                handler,
            ))
        }
    }

    handler.Init(caches.Session())
}

// RegisterAdapter registers adapter F for event E
func (r *Router) RegisterAdapter(e Event, f AdapterFunc) {
    r.Lock()
    defer r.Unlock()

    _, ok := r.EventHandlers[e]
    if !ok {
        r.EventHandlers[e] = []AdapterFunc{}
    }

    r.EventHandlers[e] = append(r.EventHandlers[e], f)
}

// Changes the active prefix handler
func (r *Router) SetPrefixHandler(h PrefixHandler) {
    r.Lock()
    defer r.Unlock()

    r.prefixHandler = h
}

// Changes the active panic handler
func (r *Router) SetPanicHandler(h PanicHandler) {
    r.Lock()
    defer r.Unlock()

    r.panicHandler = h
}

// Changes the active error handler
func (r *Router) SetParseErrorHandler(h ParseErrorHandler) {
    r.Lock()
    defer r.Unlock()

    r.parseErrorHandler = h
}

// Changes the active last resort
func (r *Router) SetLastResort(f POGOFuncW1) {
    r.Lock()
    defer r.Unlock()

    r.lastResort = f
}
