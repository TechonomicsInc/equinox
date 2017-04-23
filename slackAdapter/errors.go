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

func PanicHandler(err interface{}, withTrace bool, args ...*wormhole.Wormhole) {
    msg := args[0].AsBox().(*slack.MessageEvent)
    rtm := equinox.GetSession().(*slack.RTM)
    trace := ""

    if withTrace {
        buf := make([]byte, 1<<16)
        stackSize := runtime.Stack(buf, false)

        trace += string(buf[0:stackSize])
        trace += "\n\n"
    }

    rtm.SendMessage(rtm.NewOutgoingMessage(
        "Error :scream:\n```\n"+fmt.Sprintf("%v%v", err, trace)+"\n```",
        msg.Channel,
    ))
}
