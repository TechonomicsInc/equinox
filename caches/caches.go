package caches

import (
    "time"
    "sync"
)

const (
    DEFAULT_CACHE_EXPIRATION = int64(15*time.Minute)
)

var (
    mutex     sync.RWMutex
    container map[string]*Item
)

func Get(id string) interface {} {
    mutex.RLock()

    item, err := container[id]
    if err {
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