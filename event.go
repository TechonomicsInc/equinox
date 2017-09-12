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

//go:generate stringer -type Event

// Event is a simple enum that is used to register and dispatch events
type Event int

const (
    // Called before any message parsing is done
    MESSAGE_PRE_ANALYZE Event = iota

    // Called after the message has been identified as "for us" but before any advanced parsing
    MESSAGE_ANALYZE

    // Called after the parsing is finished
    MESSAGE_POST_ANALYZE

    // Called before the matched handler is executed
    HANDLER_PRE_EXECUTE

    // Called after the matched handler was executed
    HANDLER_POST_EXECUTE

    // Called before executing the last resort
    LAST_RESORT_PRE_EXECUTE

    // Called after executing the last resort
    LAST_RESORT_POST_EXECUTE

    // Called during message analysis when a @mention for us is encountered
    MENTION_FOUND

    // Called when the bot is mentioned in a message that doesn't look like a command
    MENTION_UNMAPPED
)
