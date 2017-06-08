package equinox

import (
    "code.lukas.moe/x/wormhole"
    "strings"
)

// PrefixHandler is used to check if  any prefix is available (eg for server based prefixes.
// If a prefix is found the function should return a wormhole to the string.
// If not a wormhole to "NO_PREFIX_FOUND" is expected
type PrefixHandler func(args ...interface{}) interface{}

// NewStaticPrefix constructs a PrefixHandler that always returns a wormhole to $prefix.
func NewStaticPrefix(prefix string) PrefixHandler {
    return func(args ...interface{}) interface{} {
        return wormhole.ToString(prefix)
    }
}

// DefaultPrefixAdapter is the default adapter that checks if the prefix is present in a message.
func DefaultPrefixAdapter(args ...interface{}) AdapterEvent {
    r := args[0].(*Router)
    p := r.prefixHandler(args[1:]...).(string)

    if p == "" {
        return STOP_EXECUTION
    }

    if strings.HasPrefix(args[1].(string), p) {
        return CONTINUE_EXECUTION
    }

    return STOP_EXECUTION
}
