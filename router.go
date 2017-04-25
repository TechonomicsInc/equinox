package equinox

import (
    "code.lukas.moe/x/wormhole"
    "reflect"
    "regexp"
    "runtime"
    "strconv"
    "strings"
    "time"
)

func (r *Router) Dispatch(e Event, args ...*wormhole.Wormhole) (ret AdapterEvent) {
    ret = NO_HANDLERS_REGISTERED

    // Check if any handlers are defined
    if _, ok := r.EventHandlers[e]; !ok {
        return
    }

    // Check if there are more than 0 defined handlers
    if len(r.EventHandlers[e]) == 0 {
        return
    }

    // Loop through the handlers
    for _, handler := range r.EventHandlers[e] {

        ret = handler(args...)

        onDebug(func() {
            handlerName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
            handlerName = strings.Replace(handlerName, "code.lukas.moe", "", -1)
            handlerName = strings.Replace(handlerName, "/x/", "", -1)
            handlerName = strings.Replace(handlerName, "equinox/", "", -1)

            logf("[DISPATCHER] %v -> %v -> %v", e.String(), handlerName, ret.String())
        })

        if ret.ShouldAbort() {
            return
        }
    }

    return
}

func (r *Router) Handle(msg string, input *wormhole.Wormhole) {
    var (
        start, end float64
        debugDefer = func() {}
    )

    onDebug(func() {
        start = float64(time.Now().UnixNano())

        debugDefer = func() {
            end = float64(time.Now().UnixNano())
            logf("[DISPATCHER] Handle() call took %f ms", (end-start)/float64(time.Millisecond))
        }
    })

    defer debugDefer()
    defer AdapterPanicHandler()

    r.Dispatch(MESSAGE_PRE_ANALYZE, wormhole.To(r), wormhole.ToString(msg), input).Act()
    r.Dispatch(MESSAGE_ANALYZE, input).Act()

    // Split message into fields
    messageFields := strings.Fields(msg)

    // Loop through all listeners
    for listener, handlers := range r.Routes {
        // Check if the listener is a RegExp
        if listener.IsRegexp {
            expr := listener.Content
            if strings.Contains(expr, "{p}") {
                expr = strings.Replace(expr, "{p}", *r.prefixHandler().AsString(), 1)
            }
            regex := regexp.MustCompile(expr)

            if regex.MatchString(msg) {
                onDebug(func() {
                    logf("[LISTENER] Triggered %s", expr)
                })

                // Extract matches using regex
                matchMap := map[string]string{}
                matches := regex.FindAllString(msg, -1)

                // Convert array to a map
                for i, match := range matches {
                    matchMap["match_"+strconv.Itoa(i)] = match
                }

                r.Dispatch(MESSAGE_POST_ANALYZE, input).Act()

                // If it's a matching regexp call all handlers
                for _, handler := range handlers {
                    go r.execHandler(
                        handler,
                        input,
                        messageFields[0],
                        strings.Join(messageFields[1:], " "),
                        matchMap,
                    )
                }

                // Kill outer loop because of the match
                return
            }
        }

        // Split the listener into fields
        listenerFields := strings.Fields(listener.Content)

        // Check if the handler expects @mentions
        if strings.Contains(listenerFields[0], "{@}") {
            // Check if any mentions are present
            r.Dispatch(MESSAGE_CHECK_MENTIONS, input).Act()

            // Check if mentions for us are present
            r.Dispatch(MESSAGE_CHECK_OUR_MENTIONS, input).Act()

            r.Dispatch(MESSAGE_POST_ANALYZE, input).Act()

            // If mentions are present call the module
            for _, handler := range handlers {
                onDebug(func() {
                    logf("[LISTENER] Triggered %s", listenerFields[0])
                })
                go r.execHandler(handler, input, messageFields[0], strings.Join(messageFields[1:], " "), nil)
            }

            // Kill outer loop
            return
        }

        // Replace the prefix-placeholder if present
        if strings.Contains(listenerFields[0], "{p}") {
            listenerFields[0] = strings.Replace(listenerFields[0], "{p}", *r.prefixHandler().AsString(), -1)
        }

        // Skip iteration if the current listener doesn't match
        if messageFields[0] != listenerFields[0] {
            continue
        }

        // Prepare container vars for command parsing
        actionParams := map[string]string{}
        i := 0

        // If the message contains less fields than the listener try send an error
        if len(messageFields) < len(listenerFields) {
            r.parseErrorHandler(
                messageFields[0],
                input,
                wormhole.ToString(
                    "Argument count mismatch.\n"+
                        strconv.Itoa(len(messageFields)-1)+ " != "+ strconv.Itoa(len(listenerFields)-1),
                ),
            )
            return
        }

        for i = 0; i < len(listenerFields); i++ {
            key := listenerFields[i]
            key = strings.Replace(key, "{", "", 1)
            key = strings.Replace(key, "}", "", 1)

            actionParams[key] = messageFields[i]
        }

        // Append unparsed text to the map
        unparsed := []string{}
        for ; i < len(messageFields); i++ {
            unparsed = append(unparsed, messageFields[i])
        }

        actionParams["unparsed"] = strings.Join(unparsed, " ")

        // Replace prefix entry
        actionParams["command"] = actionParams[messageFields[0]]
        delete(actionParams, messageFields[0])

        r.Dispatch(MESSAGE_POST_ANALYZE, input).Act()

        // Call handlers
        for _, handler := range handlers {
            onDebug(func() {
                logf("[LISTENER] Triggered %s", listenerFields[0])
            })
            go r.execHandler(handler, input, messageFields[0], strings.Join(messageFields[1:], " "), actionParams)
        }

        // Exit outer loop
        return
    }
}

func (r *Router) execHandler(
    handler Handler,
    input *wormhole.Wormhole,
    command string,
    content string,
    actionParams map[string]string,
) {
    // Defer post-execute handlers
    defer func() {
        e := recover()
        if e != nil {
            r.panicHandler(e, r.debugMode, input)
            return
        }

        r.Dispatch(HANDLER_POST_EXECUTE, input)
    }()

    // Call pre execute handlers
    if r.Dispatch(HANDLER_PRE_EXECUTE, input).ShouldAbort() {
        return
    }

    // Call action
    handler.Action(command, content, actionParams, input)
}
