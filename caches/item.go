package caches

import "time"

const (
    NO_TIMEOUT               = -1
    DEFAULT_CACHE_EXPIRATION = int64(1 * time.Minute)
    MAX_AGE                  = int64(2 * time.Minute)
)

type Item struct {
    Content    interface{}
    Creation   int64
    LastAccess int64
    Timeout    int64
}

func NewItem(content interface{}) *Item {
    return &Item{
        Content:    content,
        Creation:   time.Now().Unix(),
        LastAccess: time.Now().Unix(),
        Timeout:    DEFAULT_CACHE_EXPIRATION,
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
    // Check if the item is an exception
    if i.Timeout == NO_TIMEOUT {
        return false
    }

    // Check if it reached the max age
    if time.Now().Unix() > i.Creation+MAX_AGE {
        return true
    }

    // Check if it timed out
    if time.Now().Unix() > i.LastAccess+i.Timeout {
        return true
    }

    return false
}
