package caches

import (
    "sync"
    "time"
    "github.com/bwmarrin/discordgo"
)

var (
    mutex     = sync.RWMutex{}
    container = map[string]*Item{}
)

func Get(id string) interface{} {
    mutex.Lock()

    item, ok := container[id]
    if !ok {
        mutex.Unlock()
        return nil
    }

    item.LastAccess = time.Now().Unix()
    mutex.Unlock()

    if item.IsExpired() {
        defer Cleanup()
    }

    return item.Content
}

func Set(id string, item *Item) {
    mutex.Lock()
    container[id] = item
    container[id].LastAccess = time.Now().Unix()
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

func Session() *discordgo.Session {
    return Get("session").(*discordgo.Session)
}

func SetSession(session *discordgo.Session) {
    Set(
        "session",
        NewItem(session).SetTimeout(NO_TIMEOUT),
    )
}
