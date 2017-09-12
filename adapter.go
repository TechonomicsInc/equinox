/*
 * Copyright (C) 2017 Subliminal Apps
 *
 * Licensed under the EUPL, Version 1.1 only (the "Licence");
 *
 * You may not use this work except in compliance with the Licence.
 * You may obtain a copy of the Licence at:
 * <https://joinup.ec.europa.eu/software/page/eupl>
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Licence is distributed on an "AS IS" basis,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the Licence for the specific language governing permissions and limitations
 * under the Licence.
 */

package equinox

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

// AdapterFunc defines a function that takes N interfaces and returns an AdapterEvent
type AdapterFunc func(msg *discordgo.Message) AdapterEvent

// AdapterPanic is thrown by adapters to indicate which Event caused the panic
type AdapterPanic struct {
    Event  AdapterEvent
    Reason interface{}
}

// AdapterPanicHandler should be deferred BEFORE executing adapters
func AdapterPanicHandler() {
    if e := recover(); e != nil {
        exc, ok := e.(AdapterPanic)

        if ok {
            OnDebug(func() {
                log("Caught AdapterPanic")
                fmt.Printf("%#v\n", exc)
            })

            if exc.Event != PANIC {
                return
            }
        }

        panic(e)
    }
}
