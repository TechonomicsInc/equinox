package equinox

import (
    "fmt"
)

// NewRouter constructs a router object with some default adapters
func NewRouter() *Router {
    r := &Router{}

    r.Lock()
    r.EventHandlers = make(map[Event][]AdapterFunc)
    r.Routes = make(map[*Listener][]Handler)
    r.Unlock()

    r.UseDebugMode(false)
    r.SetPanicHandler(DefaultPanicHandler)
    r.RegisterAdapter(MESSAGE_PRE_ANALYZE, DefaultPrefixAdapter)

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

    for _, l := range (handler).Listeners() {
        if _, ok := r.Routes[l]; !ok {
            r.Routes[l] = append(r.Routes[l], handler)
        } else {
            panic(fmt.Errorf(
                "Tried to add duplicate route %s with handler \n%#v",
                l,
                handler,
            ))
        }
    }

    handler.Init()
}

func (r *Router) RegisterAdapter(e Event, f AdapterFunc) {
    r.Lock()
    defer r.Unlock()

    _, ok := r.EventHandlers[e]
    if !ok {
        r.EventHandlers[e] = []AdapterFunc{}
    }

    r.EventHandlers[e] = append(r.EventHandlers[e], f)
}

func (r *Router) SetPrefixHandler(h PrefixHandler) {
    r.Lock()
    defer r.Unlock()

    r.prefixHandler = h
}

func (r *Router) SetPanicHandler(h PanicHandler) {
    r.Lock()
    defer r.Unlock()

    r.panicHandler = h
}

func (r *Router) SetParseErrorHandler(h ParseErrorHandler) {
    r.Lock()
    defer r.Unlock()

    r.parseErrorHandler = h
}

func (r *Router) SetLastResort(f POGOFunc) {
    r.Lock()
    defer r.Unlock()

    r.lastResort = f
}
