/*
 * Copyright (C) 2017 Subliminal Apps
 *
 * Licensed under the EUPL, Version 1.1 only (the "Licence");
 *
 * You may not use this work except in compliance with the Licence.
 * You may obtain a copy of the Licence at:
 * <https://joinup.ec.europa.eu/software/page/eupl>
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Licence is distributed on an "AS IS" basis,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the Licence for the specific language governing permissions and limitations
 * under the Licence.
 */

package equinox

import (
    "fmt"
    "strings"
    "code.lukas.moe/x/equinox/caches"
    "regexp"
    "code.lukas.moe/x/equinox/annotations"
)

// NewRouter constructs a router object with some default adapters
func NewRouter() *Router {
    r := &Router{}

    r.Lock()
    r.EventHandlers = make(map[Event][]AdapterFunc)
    r.Routes = make(map[string]Handler)
    r.RuntimeAdapters = make(map[Handler][]RuntimeAdapter)
    r.AnnotationHandlers = make(map[string][]AnnotationHandler)
    r.Unlock()

    r.UseDebugMode(false)
    r.IgnoreCommandCase(false)
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

    meta := annotations.Parse(handler.Meta())
    for _, annotation := range meta {
        OnDebug(func() {
            logf("--- [ANNOTATION] [%s] -> %v", annotation.Key, annotation.Value)
        })

        switch annotation.Key {
        case "PrefixListeners":
            for _, listener := range annotation.Value {
                r.addListener("{p}"+listener, handler)
            }

        case "MentionListeners":
            for _, listener := range annotation.Value {
                r.addListener("{@}"+listener, handler)
            }

        case "Listeners":
            paramExpression := regexp.MustCompile(`{(.*?)}`)

            for _, l := range annotation.Value {
                parts := strings.Fields(l)
                l = parts[0]

                // Perform parameter expansion if needed
                lExprs := strings.Split(
                    paramExpression.FindAllStringSubmatch(l, -1)[0][1],
                    ",",
                )
                l = paramExpression.ReplaceAllString(l, "")

                if len(parts) > 1 {
                    parts = parts[1:]
                } else {
                    parts = []string{}
                }

                for _, lExpr := range lExprs {
                    r.addListener("{"+lExpr+"}"+l, handler)
                }
            }

        default:
            if handlers, ok := r.AnnotationHandlers[annotation.Key]; ok {
                for _, h := range handlers {
                    h(annotation, handler, r)
                }
            }
        }
    }

    handler.Init(caches.Session())
}

func (r *Router) addListener(listener string, handler Handler) {
    OnDebug(func() {
        logf("------ Found listener: %s", listener)
    })

    if _, ok := r.Routes[listener]; !ok {
        r.Routes[listener] = handler
    } else {
        panic(fmt.Errorf(
            "Tried to add duplicate route %s with handler \n%#v",
            listener,
            handler,
        ))
    }
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

func (r *Router) RegisterAnnotationHandler(annotation string, f AnnotationHandler) {
    r.Lock()
    defer r.Unlock()

    _, ok := r.AnnotationHandlers[annotation]
    if !ok {
        r.AnnotationHandlers[annotation] = []AnnotationHandler{}
    }

    r.AnnotationHandlers[annotation] = append(r.AnnotationHandlers[annotation], f)
}

func (r *Router) RegisterRuntimeAdapter(handler Handler, f RuntimeAdapter) {
    r.Lock()
    defer r.Unlock()

    _, ok := r.RuntimeAdapters[handler]
    if !ok {
        r.RuntimeAdapters[handler] = []RuntimeAdapter{}
    }

    r.RuntimeAdapters[handler] = append(r.RuntimeAdapters[handler], f)
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

func (r *Router) IgnoreCommandCase(v bool) {
    r.ignoreCommandCase = v
}
