package caches

import (
    "sync"
    "time"
)

const (
    DEFAULT_CACHE_EXPIRATION = int64(15 * time.Minute)
)

var (
    mutex     = sync.RWMutex{}
    container = map[string]*Item{}
)

func Get(id string) interface{} {
    mutex.RLock()

    item, ok := container[id]
    if !ok {
        return nil
    }

    if item.IsExpired() {
        defer Cleanup()
    }

    defer mutex.RUnlock()
    return item.Content
}

func Set(id string, item *Item) {
    mutex.Lock()
    container[id] = item
    mutex.Unlock()
}

func Cleanup() {
    mutex.Lock()
    defer mutex.Unlock()

    for key, item := range container {
        if !item.IsExpired() {
            continue
        }

        delete(container, key)
    }
}

func Session() interface{} {
    return Get("session")
}

func SetSession(session interface{}) {
    Set(
        "session",
        NewItem(session).SetTimeout(NO_TIMEOUT),
    )
}
