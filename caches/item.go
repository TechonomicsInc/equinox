package caches

import "time"

const (
    NO_TIMEOUT = -1
    DEFAULT_CACHE_EXPIRATION = int64(15 * time.Minute)
)

type Item struct {
    Content    interface{}
    LastAccess int64
    Timeout    int64
}

func NewItem(content interface{}) *Item {
    return &Item{
        Content: content,
        LastAccess: time.Now().Unix(),
        Timeout: DEFAULT_CACHE_EXPIRATION,
    }
}

func (i *Item) SetTimeout(time int64) *Item {
    i.Timeout = time

    return i
}

func (i *Item) SetContent(content interface{}) *Item {
    i.Content = content

    return i
}

func (i *Item) IsExpired() bool {
    if i.Timeout == NO_TIMEOUT {
        return false
    }

    return time.Now().Unix() > i.LastAccess + i.Timeout
}
