package equinox

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

// ParseErrorHandler defines the signature of a func that may catch parser errors
type ParseErrorHandler func(command string, msg *discordgo.Message, err interface{})

// DefaultParseErrorHandler is the default parse error handler (prints to STDOUT)
func DefaultParseErrorHandler(command string, msg *discordgo.Message, err interface{}) {
    fmt.Printf(
        "Error while parsing %s for %#v\n%#v",
        command,
        msg,
        err,
    )
}
