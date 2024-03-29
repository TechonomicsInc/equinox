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
    "reflect"
    "runtime"
    "strings"
    "time"
    "code.lukas.moe/x/equinox/caches"
    "code.lukas.moe/x/discordgo"
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
    if len(input.Mentions) > 0 && strings.Contains(input.Content, ourMention) {
        r.Dispatch(MENTION_FOUND, input)

        // If it looks like a command, search for a command.
        // Otherwise call the last resort.
        if strings.HasPrefix(input.Content, ourMention) {
            // Dissect the message
            parts := strings.Fields(input.Content)
            content := strings.Join(parts[2:], " ")
            cmd := parts[1]

            // Ignore case of command
            if r.ignoreCommandCase {
                cmd = strings.ToLower(cmd)
            }

            // Check if a handler for this is present
            handler, ok := r.Routes["{@}"+cmd]
            if !ok {
                r.Dispatch(LAST_RESORT_PRE_EXECUTE, input).Act()
                r.lastResort(input)
                r.Dispatch(LAST_RESORT_POST_EXECUTE, input).Act()
                return
            }

            OnDebug(func() {
                logf("[HANDLE  ][%s][COMMAND] '%s'", input.ID, cmd)
                logf("[HANDLE  ][%s][CONTENT] '%s'", input.ID, content)
            })

            go r.execHandler(handler, input, cmd, content)
            return
        }

        // We have been mentioned but the message was neither a command
        // nor a pre-mention that would've triggered the last resort.
        // Dispatch an event just in case somebody needs that.
        if r.Dispatch(MENTION_UNMAPPED, input).ShouldAbort() {
            return
        }
    }

    // Do some message analyzing checks
    // For example: Check if the message contains the prefix
    r.Dispatch(MESSAGE_ANALYZE, input).Act()

    // Dissect the message
    cmd := ""
    content := ""
    parts := strings.Fields(input.Content)
    prefix := r.prefixHandler(input)

    if prefix != "" {
        cmd = strings.Replace(parts[0], prefix, "", 1)
        content = strings.Join(parts[1:], " ")
    }

    // Ignore case of command if needed
    if r.ignoreCommandCase {
        cmd = strings.ToLower(cmd)
    }

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
    go r.execHandler(handler, input, cmd, content)
}

// execHandler safely executes the passed handler and catches any possible panics
func (r *Router) execHandler(
    handler Handler,
    input *discordgo.Message,
    command string,
    content string,
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

    if adapters, ok := r.RuntimeAdapters[handler]; ok {
        for _, adapter := range adapters {
            if adapter(handler, input, r).ShouldAbort() {
                return
            }
        }
    }

    // Call action
    handler.Action(command, content, input, caches.Session())
}
