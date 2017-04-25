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

func UseDebugMode(i bool) {
    useDebugMode(i)
}

func (r *Router) UseDebugMode(i bool) {
    useDebugMode(i)
}

func GetDebugMode() bool {
    return getDebugMode()
}

func (r *Router) GetDebugMode() bool {
    return getDebugMode()
}

func onDebug(fun func()) {
    if GetDebugMode() {
        fun()
    }
}
