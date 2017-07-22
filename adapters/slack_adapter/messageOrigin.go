package slack_adapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/nlopes/slack"
    "code.lukas.moe/x/equinox/caches"
)

func IgnorePrivateMessages(args ...interface{}) equinox.AdapterEvent {
    input := args[0].(*slack.MessageEvent)
    session := caches.Session().(*slack.RTM)

    channel, err := session.GetChannelInfo(input.Channel)
    if err != nil {
        return equinox.CONTINUE_EXECUTION
    }

    if !channel.IsMember {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}

func IgnoreChannelMessages(args ...interface{}) equinox.AdapterEvent {
    ret := IgnorePrivateMessages(args...)

    if ret == equinox.STOP_EXECUTION {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}
