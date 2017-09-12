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
