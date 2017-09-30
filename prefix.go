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
    "code.lukas.moe/x/discordgo"
    "strings"
)

// PrefixHandler is used to get the current prefix.
// If a prefix is found the function should return the string.
// If not an empty string is expected
type PrefixHandler func(msg *discordgo.Message) string

// PrefixAdapter is a special adapter that is expected to call the PrefixHandler.
// It should then use the received information to determine whether *msg contains the prefix or not.
type PrefixAdapter func(r *Router, msg *discordgo.Message) AdapterEvent

// NewStaticPrefix constructs a PrefixHandler that always returns $prefix.
// Useful for simple or small bots.
func NewStaticPrefix(prefix string) PrefixHandler {
    return func(msg *discordgo.Message) string {
        return prefix
    }
}

// DefaultPrefixAdapter is the default adapter that checks if the prefix is present in a message.
func DefaultPrefixAdapter(r *Router, msg *discordgo.Message) AdapterEvent {
    p := r.prefixHandler(msg)

    if p == "" || (len(msg.Content) < len(p)+1) {
        return STOP_EXECUTION
    }

    if strings.HasPrefix(msg.Content, p) {
        return CONTINUE_EXECUTION
    }

    return STOP_EXECUTION
}
