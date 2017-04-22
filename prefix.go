package equinox

import (
    "code.lukas.moe/x/wormhole"
    "strings"
)

// PrefixHandler is used to check if  any prefix is available (eg for server based prefixes.
// If a prefix is found the function should return a wormhole to the string.
// If not a wormhole to "NO_PREFIX_FOUND" is expected
type PrefixHandler func(args ...*wormhole.Wormhole) *wormhole.Wormhole

// NewStaticPrefix constructs a PrefixHandler that always returns a wormhole to $prefix.
func NewStaticPrefix(prefix string) PrefixHandler {
    return func(args ...*wormhole.Wormhole) *wormhole.Wormhole {
        return wormhole.ToString(prefix)
    }
}

// DefaultPrefixAdapter is the default adapter that checks if the prefix is present in a message.
func DefaultPrefixAdapter(args ...*wormhole.Wormhole) AdapterEvent {
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
