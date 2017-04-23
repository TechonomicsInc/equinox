package slackAdapter

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/wormhole"
    "fmt"
    "github.com/nlopes/slack"
    "runtime"
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

func PanicHandler(err interface{}, args ...*wormhole.Wormhole) {
    msg := args[0].AsBox().(*slack.MessageEvent)
    rtm := equinox.GetSession().(*slack.RTM)

    buf := make([]byte, 1<<16)
    stackSize := runtime.Stack(buf, false)

    rtm.SendMessage(rtm.NewOutgoingMessage(
        "Error :scream:\n```\n" + fmt.Sprintf("%v\n%v", err, string(buf[0:stackSize])) + "\n```",
        msg.Channel,
    ))
}
