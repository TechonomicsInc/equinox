package slack_adapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/nlopes/slack"
    "code.lukas.moe/x/equinox/caches"
)

func StartTyping(args ...interface{}) equinox.AdapterEvent {
    session := caches.Get(caches.SESSION).(*slack.RTM)

    session.SendMessage(session.NewTypingMessage(
        args[0].(*slack.MessageEvent).Channel,
    ))

    return equinox.CONTINUE_EXECUTION
}
