package equinox

import (
    "code.lukas.moe/x/wormhole"
    "fmt"
    "reflect"
    "regexp"
    "runtime"
    "strconv"
    "strings"
    "time"
)

func (r *Router) Dispatch(e Event, args ...*wormhole.Wormhole) (ret AdapterEvent) {
    defer PanicAdapter()

    // If there are no handlers for the event just return
    if _, ok := r.EventHandlers[e]; !ok {
        return
    }

    if len(r.EventHandlers[e]) == 0 {
        return
    }

    // Loop through the handlers
    for _, handler := range r.EventHandlers[e] {

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

    return
}

func (r *Router) Handle(msg string, input *wormhole.Wormhole) {
    // Deferred counter for debugging
    start := time.Now().UnixNano()
    defer func() {
        end := time.Now().UnixNano()
        duration := time.Duration(end - start)

        fmt.Printf("[DISPATCHER] Handle() call took %f ms\n", float64(duration)/float64(time.Millisecond))
    }()

    // Init adapter var
    ret := CONTINUE_EXECUTION

    // Call PRE_ANALYZE adapters
    ret = r.Dispatch(MESSAGE_PRE_ANALYZE, wormhole.To(r), wormhole.ToString(msg), input)
    if ret == STOP_EXECUTION {
        return
    }

    // Call ANALYZE adapters
    ret = r.Dispatch(MESSAGE_ANALYZE, input)
    if ret == STOP_EXECUTION {
        return
    }

    // Split message into fields
    messageFields := strings.Fields(msg)

    // Loop through all listeners
    for listener, handlers := range r.Routes {
        // Check if the listener is a RegExp
        if listener.IsRegexp && listener.Content.(*regexp.Regexp).MatchString(msg) {
            // If it's a matching regexp call all handlers
            for _, handler := range handlers {
                go r.execHandler(handler, input, messageFields[0], strings.Join(messageFields[1:], " "), nil)
            }

            // Break outer loop because of the match
            break
        }

        // Split the listener into fields
        listenerFields := strings.Fields(listener.Content.(string))

        // Check if the handler expects @mentions
        if strings.Contains(listenerFields[0], "{@}") {
            // Call @mention adapters
            ret = r.Dispatch(MESSAGE_CHECK_OUR_MENTIONS, input)
            if ret == STOP_EXECUTION {
                return
            }

            // If mentions are present call the module
            for _, handler := range handlers {
                go r.execHandler(handler, input, messageFields[0], strings.Join(messageFields[1:], " "), nil)
            }

            // Break outer loop
            break
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
            break
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

        // Call handlers
        for _, handler := range handlers {
            go r.execHandler(handler, input, messageFields[0], strings.Join(messageFields[1:], " "), actionParams)
        }

        // Break outer loop
        break
    }

    ret = r.Dispatch(MESSAGE_POST_ANALYZE, input)
}

func (r *Router) execHandler(
    handler *Handler,
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

        ret := r.Dispatch(HANDLER_POST_EXECUTE, input)
        if ret == STOP_EXECUTION {
            return
        }
    }()

    // Call pre execute handlers
    ret := r.Dispatch(HANDLER_PRE_EXECUTE, input)
    if ret == STOP_EXECUTION {
        return
    }

    // Call action
    (*handler).Action(command, content, actionParams, input)
}
