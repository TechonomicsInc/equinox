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

//go:generate stringer -type AdapterEvent

type AdapterEvent int

const (
    // Indicates that no handlers were registered.
    // It's recommended to act like the return value was CONTINUE_EXECUTION if this happens.
    NO_HANDLERS_REGISTERED AdapterEvent = iota

    // Indicates that the outer routine should continue to work normally
    CONTINUE_EXECUTION

    // Indicates that the outer routine should return as soon as possible
    STOP_EXECUTION

    // Indicates that the outer routine should throw a panic
    PANIC
)

// ShouldAbort is a convenience function that checks if a state was encountered that
// requires the outer routine to exit.
func (a AdapterEvent) ShouldAbort() bool {
    return a == STOP_EXECUTION || a == PANIC
}

// Act is a convenience function checks all possible event values and acts accordingly for you.
// That means:
// - CONTINUE_EXECUTION and NO_HANDLERS_REGISTERED is a NOOP.
// - Everything else panics with a proper (and parsable) error message.
func (a AdapterEvent) Act() {
    switch a {
    case
        CONTINUE_EXECUTION,
        NO_HANDLERS_REGISTERED:
        return

    case
        STOP_EXECUTION:
        panic(AdapterPanic{a, "NORMAL_EVENT_TERMINATION"})

    case PANIC:
        panic(AdapterPanic{a, "EXPECTED_ABNORMAL_EVENT_TERMINATION"})

    default:
        panic(AdapterPanic{a, "UNEXPECTED_ABNORMAL_EVENT_TERMINATION"})
    }
}
