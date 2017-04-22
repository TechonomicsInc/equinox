package equinox

import (
    "code.lukas.moe/x/wormhole"
    "fmt"
)

type PanicHandler func(input *wormhole.Wormhole)

func DefaultPanicHandler(input *wormhole.Wormhole) {
    err := recover()
    if err != nil {
        fmt.Printf("\n\nFailure encountered.\n\nHint:\n%#v\n\nActual Error:\n%#v\n\n", input.AsBox(), err)
        panic(err)
    }
}
