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
    mutex.RLock()
    defer TouchItem(id)
    defer mutex.RUnlock()

    item, ok := container[id]
    if !ok {
        return nil
    }

    return item.Content
}

func Set(id string, item *Item) {
    mutex.Lock()
    defer TouchItem(id)
    defer mutex.Unlock()

    container[id] = item
}

func TouchItem(id string) {
    mutex.Lock()
    defer mutex.Unlock()

    if item, ok := container[id]; ok {
        item.LastAccess = time.Now().Unix()
    }
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

func Manage() {
    for {
        time.Sleep(5 * time.Second)
        Cleanup()
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
