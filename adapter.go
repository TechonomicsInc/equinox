package equinox

import "code.lukas.moe/x/wormhole"

// AdapterEvent is an enum that tells the router to continue/stop execution
type AdapterEvent int

const (
    CONTINUE_EXECUTION AdapterEvent = iota
    STOP_EXECUTION
    PANIC
)

// AdapterFunc is a function that takes N wormholes and returns an AdapterEvent
type AdapterFunc func(args ...*wormhole.Wormhole) AdapterEvent

// POGOFunc (Plain old Go Func) is a simple void func without parameters
type POGOFunc func()

// NOOP is a small helper for router parameters that require POGOFunc's
func NOOP() {}

// AdapterPanic is thrown by adapters to indicate which Event caused the panic
type AdapterPanic struct {
    Event  AdapterEvent
    Reason interface{}
}

// PanicAdapter should be deferred when executing adapters
func PanicAdapter() {
    if e := recover(); e != nil {
        exc, ok := e.(AdapterPanic)

        if ok && exc.Event != PANIC {
            return
        }

        panic(e)
    }
}
