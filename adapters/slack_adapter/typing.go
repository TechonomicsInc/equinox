package slack_adapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/nlopes/slack"
)

func StartTyping(args ...interface{}) equinox.AdapterEvent {
    session := equinox.GetSession().(*slack.RTM)

    session.SendMessage(session.NewTypingMessage(
        args[0].(*slack.MessageEvent).Channel,
    ))

    return equinox.CONTINUE_EXECUTION
}
