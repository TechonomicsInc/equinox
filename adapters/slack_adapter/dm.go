package slack_adapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/nlopes/slack"
    "code.lukas.moe/x/equinox/caches"
)

func IsDMOrPrefixed(args ...interface{}) equinox.AdapterEvent {
    if IsDM(args...) == equinox.CONTINUE_EXECUTION {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.DefaultPrefixAdapter(args...)
}

func IsDM(args ...interface{}) equinox.AdapterEvent {
    rtm := caches.Get(caches.SESSION).(*slack.RTM)
    m := args[0].(*slack.MessageEvent)

    _, _, ch, e := rtm.OpenIMChannel(m.User)
    if e != nil {
        return equinox.STOP_EXECUTION
    }

    if m.Channel == ch {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}
