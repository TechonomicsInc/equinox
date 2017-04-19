package slackAdapter

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/wormhole"
    "github.com/nlopes/slack"
)

func IgnorePrivateMessages(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    input := args[0].AsBox().(*slack.MessageEvent)
    session := equinox.GetSession().AsBox().(*slack.RTM)

    channel, err := session.GetChannelInfo(input.Channel)
    if err != nil {
        return equinox.CONTINUE_EXECUTION
    }

    if !channel.IsMember {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}

func IgnoreChannelMessages(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    ret := IgnorePrivateMessages(args...)

    if ret == equinox.STOP_EXECUTION {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}
