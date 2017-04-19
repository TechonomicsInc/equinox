package equinox

import "code.lukas.moe/x/wormhole"

type AdapterEvent int

type AdapterFunc func(args ...*wormhole.Wormhole) AdapterEvent

type POGOFunc func()

func NOOP() {}

type AdapterPanic struct {
    Event  AdapterEvent
    Reason string
}

func PanicAdapter() {
    if e := recover(); e != nil {
        exc, ok := e.(AdapterPanic)

        if ok && exc.Event != PANIC {
            return
        }

        panic(e)
    }
}

const (
    CONTINUE_EXECUTION AdapterEvent = iota
    STOP_EXECUTION
    PANIC
)
