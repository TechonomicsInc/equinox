package equinox

import (
    "fmt"
)

// ParseErrorHandler defines the signature of a func that may catch parser errors
type ParseErrorHandler func(command string, msg interface{}, err interface{})

// DefaultParseErrorHandler is the default parse error handler (prints to STDOUT)
func DefaultParseErrorHandler(command string, msg interface{}, err interface{}) {
    fmt.Printf(
        "Error while parsing %s for %#v\n%#v",
        command,
        msg,
        err,
    )
}
