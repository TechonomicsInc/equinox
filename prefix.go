package equinox

import (
    "strings"
    "github.com/bwmarrin/discordgo"
)

// PrefixHandler is used to check if  any prefix is available (eg for server based prefixes.
// If a prefix is found the function should return the string.
// If not an empty string is expected
type PrefixHandler func(msg *discordgo.Message) string

// NewStaticPrefix constructs a PrefixHandler that always returns $prefix.
func NewStaticPrefix(prefix string) PrefixHandler {
    return func(msg *discordgo.Message) string {
        return prefix
    }
}

// DefaultPrefixAdapter is the default adapter that checks if the prefix is present in a message.
func DefaultPrefixAdapter(r *Router, msg *discordgo.Message) AdapterEvent {
    p := r.prefixHandler(msg)

    if p == "" {
        return STOP_EXECUTION
    }

    if strings.HasPrefix(msg.Content, p) {
        return CONTINUE_EXECUTION
    }

    return STOP_EXECUTION
}
