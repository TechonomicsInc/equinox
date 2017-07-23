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
    Listeners() []string

    // Called if any of the listeners matched
    Action(
        command string,
        content string,
        params map[string]string,
        msg *discordgo.Message,
        session *discordgo.Session,
    )
}

type ListenerMeta struct {
    Expression []string
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

    // map of command -> handlers
    Routes map[string]Handler

    // map of command -> expression
    RouteMeta map[string]ListenerMeta
}
