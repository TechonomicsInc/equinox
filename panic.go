package equinox

import (
    "fmt"
)

// PanicHandler defines a function that the router uses to handle panics.
// Bots may override this to send messages to the chat service or SAAS-apps like sentry.io
type PanicHandler func(err interface{}, withTrace bool, args ...interface{})

// DefaultPanicHandler is the simplest implementation of a PanicHandler (prints to STDOUT and optionally panics)
func DefaultPanicHandler(err interface{}, withTrace bool, args ...interface{}) {
    fmt.Printf("\n\nFailure encountered.\n\nHint:\n%#v\n\nActual Error:\n%#v\n\n", args[0], err)

    if withTrace {
        panic(err)
    }
}
