package equinox

import (
    "code.lukas.moe/x/wormhole"
    "fmt"
)

// PanicHandler defines a function that the router uses to handle panics.
// Bots may override this to send messages to the chat service or SAAS-apps like sentry.io
type PanicHandler func(args ...*wormhole.Wormhole)

// DefaultPanicHandler is the simplest implementation of a PanicHandler
func DefaultPanicHandler(args ...*wormhole.Wormhole) {
    err := recover()

    if err != nil {
        fmt.Printf("\n\nFailure encountered.\n\nHint:\n%#v\n\nActual Error:\n%#v\n\n", args[0].AsBox(), err)
        panic(err)
    }
}
