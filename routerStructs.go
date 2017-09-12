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
    "sync"
    "github.com/bwmarrin/discordgo"
)

// Handler defines the basic layout of Handler structs
type Handler interface {
    // Called after the handler is registered through AddRoute()
    Init(session *discordgo.Session)

    // Called to retrieve the Listener patterns
    Meta() string

    // Called if any of the listeners matched
    Action(
        command string,
        content string,
        msg *discordgo.Message,
        session *discordgo.Session,
    )
}

// TODO: write docs about the most important thing in equinox ._.
type Router struct {
    sync.RWMutex

    debugMode         bool
    ignoreCommandCase bool

    lastResort        POGOFuncW1
    prefixHandler     PrefixHandler
    panicHandler      PanicHandler
    parseErrorHandler ParseErrorHandler

    EventHandlers map[Event][]AdapterFunc

    // map of command -> handlers
    Routes map[string]Handler

    // Functions that help to give unknown annotations a purpose
    AnnotationHandlers map[string][]AnnotationHandler

    // Works exactly like classic adapters but is bound to a handler
    // instead of an equinox event.
    RuntimeAdapters map[Handler][]RuntimeAdapter
}
