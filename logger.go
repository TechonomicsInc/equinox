package equinox

import (
    "fmt"
    "time"
)

func (r *Router) logf(format string, a ...interface{}) {
    r.log(fmt.Sprintf(format, a...))
}

func (r *Router) log(msg string) {
    if !r.debugMode {
        return
    }

    fmt.Printf(
        "[%s] (%s) %s\n",
        time.Now().Format("15:04:05 02-01-2006"),
        "TRACKY-DEBUG",
        msg,
    )
}
