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

import "sync"

var (
    debugMode  = false
    debugMutex sync.RWMutex
)

func useDebugMode(i bool) {
    debugMutex.Lock()
    defer debugMutex.Unlock()

    debugMode = i
}

func getDebugMode() bool {
    debugMutex.RLock()
    defer debugMutex.RUnlock()

    return debugMode
}

// UseDebugMode enables/disables the equinox debugging
func UseDebugMode(i bool) {
    useDebugMode(i)
}

// UseDebugMode enables/disables the equinox debugging.
// This is a convenience function bound to the Router struct.
// The effect will be global.
func (r *Router) UseDebugMode(i bool) {
    useDebugMode(i)
}

// GetDebugMode returns the current debugging mode (on/off)
func GetDebugMode() bool {
    return getDebugMode()
}

// GetDebugMode returns the current debugging mode (on/off)
// This is a convenience function bound to the Router struct.
// The effect will be global.
func (r *Router) GetDebugMode() bool {
    return getDebugMode()
}

// OnDebug is a conveniece function that conditionally executes code.
// The callback will do nothing if debugging is disabled.
func OnDebug(fun func()) {
    if GetDebugMode() {
        fun()
    }
}
