package equinox

import (
    "code.lukas.moe/x/wormhole"
    "fmt"
)

type ParseErrorHandler func(command string, msg *wormhole.Wormhole, err *wormhole.Wormhole)

func DefaultParseErrorHandler(command string, msg *wormhole.Wormhole, err *wormhole.Wormhole) {
    fmt.Printf(
        "Error while parsing %s for %#v\n%#v",
        command,
        msg.AsBox(),
        err.AsBox(),
    )
}
