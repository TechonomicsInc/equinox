package equinox

import (
    "code.lukas.moe/x/wormhole"
    "github.com/davecgh/go-spew/spew"
)

// AdapterFunc is a function that takes N wormholes and returns an AdapterEvent
type AdapterFunc func(args ...*wormhole.Wormhole) AdapterEvent

// AdapterPanic is thrown by adapters to indicate which Event caused the panic
type AdapterPanic struct {
    Event  AdapterEvent
    Reason interface{}
}

// PanicAdapter should be deferred when executing adapters
func AdapterPanicHandler() {
    if e := recover(); e != nil {
        exc, ok := e.(AdapterPanic)

        if ok {
            onDebug(func() {
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

// POGOFunc (Plain old Go Func) is a simple void func without parameters
type POGOFunc func()

// NOOP is a small helper for router parameters that require POGOFunc's
func NOOP() {}
