package slackAdapter

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/wormhole"
    "github.com/nlopes/slack"
)

func StartTyping(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    equinox.GetSession().(*slack.RTM).NewTypingMessage(
        args[0].AsBox().(*slack.MessageEvent).Channel,
    )

    return equinox.CONTINUE_EXECUTION
}
