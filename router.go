package equinox

import (
    "code.lukas.moe/x/wormhole"
    "fmt"
    "reflect"
    "regexp"
    "runtime"
    "strings"
    "sync"
    "time"
)

type Handler interface {
    Init(session *wormhole.Wormhole)

    Listeners() []*Listener

    Action(
        command string,
        content string,
        params map[string]string,
        msg *wormhole.Wormhole,
    )
}

type Listener struct {
    IsRegexp bool
    Content  interface{}
}

type Router struct {
    sync.RWMutex

    lastResort    POGOFunc
    prefixHandler PrefixHandler
    panicHandler  PanicHandler

    EventHandlers map[Event][]AdapterFunc
    Routes        map[*Listener][]*Handler
}

func NewRouter() *Router {
    r := &Router{}

    r.Lock()
    r.EventHandlers = make(map[Event][]AdapterFunc)
    r.Routes = make(map[*Listener][]*Handler)
    r.Unlock()

    r.RegisterAdapter(MESSAGE_PRE_ANALYZE, PrefixAdapter)

    return r
}

func (r *Router) AddRoute(handler *Handler) {
    r.Lock()
    defer r.Unlock()

    for _, l := range (*handler).Listeners() {
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
}

func (r *Router) AddRoutes(handlers []Handler) {
    for _, h := range handlers {
        r.AddRoute(&h)
    }
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

func (r *Router) Dispatch(e Event, args ...*wormhole.Wormhole) (ret AdapterEvent) {
    defer PanicAdapter()

    for event, handlers := range r.EventHandlers {
        if event != e || handlers == nil || len(handlers) == 0 {
            continue
        }

        for _, handler := range handlers {
            ret = handler(args...)

            fmt.Printf(
                "[DISPATCHER] EQUINOX_EVENT:%v -> %v -> ADAPTER_EVENT:%v\n",
                e,
                runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name(),
                ret,
            )

            switch ret {
            case CONTINUE_EXECUTION:
                continue

            case STOP_EXECUTION:
                return
            }
        }
    }

    return
}

func (r *Router) Handle(msg string, input *wormhole.Wormhole) {
    start := time.Now().UnixNano()
    defer func() {
        end := time.Now().UnixNano()
        duration := time.Duration(end - start)

        fmt.Printf("[DISPATCHER] Handle() call took %f ms\n", float64(duration)/float64(time.Millisecond))
    }()

    var ret AdapterEvent

    ret = r.Dispatch(MESSAGE_PRE_ANALYZE, wormhole.To(r), wormhole.ToString(msg), input)
    if ret == STOP_EXECUTION {
        return
    }

    ret = r.Dispatch(MESSAGE_ANALYZE, input)
    if ret == STOP_EXECUTION {
        return
    }

    msgF := strings.Fields(msg)

    for listener, handlers := range r.Routes {
        if listener.IsRegexp && listener.Content.(*regexp.Regexp).MatchString(msg) {
            for _, handler := range handlers {
                go func() {
                    defer r.panicHandler(input)

                    (*handler).Action(msgF[0], strings.Join(msgF[1:], " "), nil, input)
                }()
            }
            break
        }
    }

    ret = r.Dispatch(MESSAGE_POST_ANALYZE, input)
}

func (r *Router) LastResort(f POGOFunc) {
    r.Lock()
    defer r.Unlock()

    r.lastResort = f
}
