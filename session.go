package equinox

import (
    "code.lukas.moe/x/wormhole"
    "errors"
    "sync"
)

var (
    sessionCache *wormhole.Wormhole
    sessionMutex sync.RWMutex
)

// SetSession saves a session in equinox
func SetSession(session *wormhole.Wormhole) {
    sessionMutex.Lock()
    defer sessionMutex.Unlock()

    sessionCache = session
}

// GetSession retrieves the session from equinox
func GetSession() interface{} {
    sessionMutex.RLock()
    defer sessionMutex.RUnlock()

    if sessionCache == nil {
        panic(errors.New("Tried to get session before equinox#SetSession() was called"))
    }

    return sessionCache.AsBox()
}
