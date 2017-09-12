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
