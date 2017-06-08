package equinox

import (
    "github.com/davecgh/go-spew/spew"
)

// AdapterFunc defines a function that takes N wormholes and returns an AdapterEvent
type AdapterFunc func(args ...interface{}) AdapterEvent

// AdapterPanic is thrown by adapters to indicate which Event caused the panic
type AdapterPanic struct {
    Event  AdapterEvent
    Reason interface{}
}

// AdapterPanicHandler should be deferred BEFORE executing adapters
func AdapterPanicHandler() {
    if e := recover(); e != nil {
        exc, ok := e.(AdapterPanic)

        if ok {
            OnDebug(func() {
                log("Caught AdapterPanic")
                spew.Dump(exc)
            })

            if exc.Event != PANIC {
                return
            }
        }

        panic(e)
    }
}
