package slack_adapter

import (
    "code.lukas.moe/x/equinox"
    "fmt"
    "github.com/nlopes/slack"
    "runtime"
    "code.lukas.moe/x/equinox/caches"
)

func ParseErrorHandler(command string, msg interface{}, err interface{}) {
    rtm := equinox.(*slack.RTM)
    msgo := msg.(*slack.MessageEvent)

    rtm.SendMessage(
        rtm.NewOutgoingMessage(
            fmt.Sprintf(
                "Parse error while processing `%s`!\n```\n%s\n```",
                command,
                err.(string),
            ),
            msgo.Channel,
        ),
    )
}

func NewPanicHandler(message string, appendix string, userCodeblock bool) equinox.PanicHandler {
    return func(err interface{}, withTrace bool, args ...interface{}) {
        msg := args[0].(*slack.MessageEvent)
        rtm := caches.Get(caches.SESSION).(*slack.RTM)
        trace := ""

        if withTrace {
            buf := make([]byte, 1<<16)
            stackSize := runtime.Stack(buf, false)

            trace += "\n\n" + string(buf[0:stackSize])
        }

        m := ""
        m += message
        if userCodeblock {
            m += "\n```\n"
        }
        m += fmt.Sprintf("%v%v", err, trace)
        if userCodeblock {
            m += "\n```\n"
        }
        m += appendix

        rtm.SendMessage(rtm.NewOutgoingMessage(m, msg.Channel, ))
    }
}

func NewDefaultPanicHandler() equinox.PanicHandler {
    return NewPanicHandler("Error :scream:", "", true)
}
