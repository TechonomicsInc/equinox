package slack_adapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/nlopes/slack"
)

func IgnoreOtherBots(args ...interface{}) equinox.AdapterEvent {
    input := args[0].(*slack.MessageEvent)

    if input.BotID != "" {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}

func IgnoreGroupMentions(args ...interface{}) equinox.AdapterEvent {
    return equinox.CONTINUE_EXECUTION
}
