package slackAdapter

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/wormhole"
    "github.com/nlopes/slack"
)

func IgnoreOtherBots(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    input := args[0].AsBox().(*slack.MessageEvent)

    if input.BotID != "" {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}

func IgnoreGroupMentions(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    return equinox.CONTINUE_EXECUTION
}
