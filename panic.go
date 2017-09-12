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

// PanicHandler defines a function that the router uses to handle panics.
// Bots may override this to send messages to the chat service or SAAS-apps like sentry.io
type PanicHandler func(err interface{}, withTrace bool, args ...interface{})

// DefaultPanicHandler is the simplest implementation of a PanicHandler (prints to STDOUT and optionally panics)
func DefaultPanicHandler(err interface{}, withTrace bool, args ...interface{}) {
    logf("\n\nFailure encountered.\n\nHint:\n%#v\n\nActual Error:\n%#v\n\n", args[0], err)

    if withTrace {
        panic(err)
    }
}
