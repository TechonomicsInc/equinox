package equinox

import (
    "code.lukas.moe/x/wormhole"
    "fmt"
)

// ParseErrorHandler defines the signature of a func that may catch parser errors
type ParseErrorHandler func(command string, msg *wormhole.Wormhole, err *wormhole.Wormhole)

// DefaultParseErrorHandler is the default parse error handler (prints to STDOUT)
func DefaultParseErrorHandler(command string, msg *wormhole.Wormhole, err *wormhole.Wormhole) {
    fmt.Printf(
        "Error while parsing %s for %#v\n%#v",
        command,
        msg.AsBox(),
        err.AsBox(),
    )
}
