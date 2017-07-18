package equinox

import (
    "fmt"
    "time"
)

func logf(format string, a ...interface{}) {
    if debugMode {
        log(fmt.Sprintf(format, a...))
    }
}

func log(msg string) {
    if debugMode {
        fmt.Printf(
            "[%s] (%s) %s\n",
            time.Now().Format("15:04:05 02-01-2006"),
            "EQUINOX",
            msg,
        )
    }
}
