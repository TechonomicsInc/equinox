package equinox

import (
    "code.lukas.moe/x/wormhole"
    "strings"
)

type PrefixHandler func(args ...*wormhole.Wormhole) *wormhole.Wormhole

func NewStaticPrefix(prefix string) PrefixHandler {
    return func(args ...*wormhole.Wormhole) *wormhole.Wormhole {
        return wormhole.ToString(prefix)
    }
}

func PrefixAdapter(args ...*wormhole.Wormhole) AdapterEvent {
    r := args[0].AsBox().(*Router)
    p := r.prefixHandler(args[1:]...).AsString()

    if *p == "" {
        return STOP_EXECUTION
    }

    if strings.HasPrefix(*args[1].AsString(), *p) {
        return CONTINUE_EXECUTION
    }

    return STOP_EXECUTION
}
