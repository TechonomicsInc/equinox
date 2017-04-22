package slackAdapter

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/wormhole"
    "fmt"
    "github.com/nlopes/slack"
)

func ParseErrorHandler(command string, msg *wormhole.Wormhole, err *wormhole.Wormhole) {
    rtm := equinox.GetSession().(*slack.RTM)
    msgo := msg.AsBox().(*slack.MessageEvent)

    rtm.SendMessage(
        rtm.NewOutgoingMessage(
            fmt.Sprintf(
                "Parse error while processing `%s`!\n```\n%s\n```",
                command,
                *err.AsString(),
            ),
            msgo.Channel,
        ),
    )
}
