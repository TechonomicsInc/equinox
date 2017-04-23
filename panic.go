package equinox

import (
    "code.lukas.moe/x/wormhole"
    "fmt"
)

// PanicHandler defines a function that the router uses to handle panics.
// Bots may override this to send messages to the chat service or SAAS-apps like sentry.io
type PanicHandler func(err interface{}, withTrace bool, args ...*wormhole.Wormhole)

// DefaultPanicHandler is the simplest implementation of a PanicHandler
func DefaultPanicHandler(err interface{}, withTrace bool, args ...*wormhole.Wormhole) {
    fmt.Printf("\n\nFailure encountered.\n\nHint:\n%#v\n\nActual Error:\n%#v\n\n", args[0].AsBox(), err)

    if withTrace {
        panic(err)
    }
}
