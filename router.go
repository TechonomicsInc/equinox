package equinox

import (
    "reflect"
    "runtime"
    "strings"
    "time"
    "code.lukas.moe/x/equinox/caches"
    "github.com/bwmarrin/discordgo"
)

// Sends an event to all registered listeners.
// If non are registered it will emit a NO_HANDLERS_REGISTERED.
func (r *Router) Dispatch(e Event, input *discordgo.Message) (ret AdapterEvent) {
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
        var (
            handlerName string
            start       float64
            end         float64
        )

        OnDebug(func() {
            handlerName = runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

            parts := strings.Split(handlerName, "/")
            handlerName = parts[len(parts)-1]

            logf("[DISPATCH][%s][EVENT ] %v -> %v", input.ID, e.String(), handlerName)
            start = float64(time.Now().UnixNano())
        })

        ret = handler(input)

        OnDebug(func() {
            end = float64(time.Now().UnixNano())

            logf("[DISPATCH][%s][RETURN] %v", input.ID, ret.String())
            logf("[DISPATCH][%s][TIMING] Call took %f ms", input.ID, (end-start)/float64(time.Millisecond))
        })

        if ret.ShouldAbort() {
            return
        }
    }

    return
}

// Handle takes an incoming message and the corresponding object.
// The message is parsed according to all registered handlers and if one of them matches
// they will be executed in the execHandler() sandbox.
func (r *Router) Handle(input *discordgo.Message) {
    var (
        start, end float64
        debugDefer = func() {}
    )

    OnDebug(func() {
        logf("[HANDLE  ][%s] Handle() was called", input.ID)

        start = float64(time.Now().UnixNano())

        debugDefer = func() {
            end = float64(time.Now().UnixNano())
            logf("[HANDLE  ][%s] Handle() call took %f ms", input.ID, (end-start)/float64(time.Millisecond))
        }
    })

    defer debugDefer()
    defer AdapterPanicHandler()

    r.Dispatch(MESSAGE_PRE_ANALYZE, input).Act()
    defer r.Dispatch(MESSAGE_POST_ANALYZE, input)

    // Check if the message contains a mention for us
    ourMention := "<@" + caches.Session().State.User.ID + ">"
    if len(input.Mentions) > 0 && strings.HasPrefix(input.Content, ourMention) {
        r.Dispatch(MENTION_FOUND, input)

        // Dissect the message
        parts := strings.Fields(input.Content)
        cmd := parts[1]
        content := strings.Join(parts[2:], " ")

        // Check if a handler for this is present
        handler, ok := r.Routes["{@}"+cmd]
        if !ok {
            r.Dispatch(LAST_RESORT_PRE_EXECUTE, input).Act()
            r.lastResort(input)
            r.Dispatch(LAST_RESORT_POST_EXECUTE, input).Act()
            return
        }

        go r.execHandler(handler, input, cmd, content, map[string]string{})
        return
    }

    // Check if the message if prefixed for us
    // First get the prefix
    prefix := r.prefixHandler(input)
    if prefix == "" {
        return
    }

    // Do some message analyzing checks
    // For example: Check if the message contains the prefix
    r.Dispatch(MESSAGE_ANALYZE, input).Act()

    // Dissect the message
    parts := strings.Fields(input.Content)
    cmd := strings.Replace(parts[0], prefix, "", 1)
    content := strings.Join(parts[1:], " ")

    OnDebug(func() {
        logf("[HANDLE  ][%s][COMMAND] '%s'", input.ID, cmd)
        logf("[HANDLE  ][%s][CONTENT] '%s'", input.ID, content)
    })

    // Check if a handler for that command is present
    handler, ok := r.Routes["{p}"+cmd]
    if !ok {
        return
    }

    // Execute
    go r.execHandler(handler, input, cmd, content, map[string]string{})
}

// execHandler safely executes the passed handler and catches any possible panics
func (r *Router) execHandler(
    handler Handler,
    input *discordgo.Message,
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
    handler.Action(command, content, actionParams, input, caches.Session())
}
