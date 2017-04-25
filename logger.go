package equinox

import (
    "fmt"
    "time"
)

func logf(format string, a ...interface{}) {
    log(fmt.Sprintf(format, a...))
}

func log(msg string) {
    fmt.Printf(
        "[%s] (%s) %s\n",
        time.Now().Format("15:04:05 02-01-2006"),
        "EQUINOX",
        msg,
    )
}
