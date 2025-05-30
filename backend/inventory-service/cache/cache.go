package cache

import (
    "sync"
    "time"
)

type Item struct {
    Value      interface{}
    Expiration int64
}

type Cache struct {
    items map[string]Item
    mu    sync.RWMutex
}

func NewCache() *Cache {
    cache := &Cache{
        items: make(map[string]Item),
    }
    go cache.startCleanupTimer()
    return cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items[key] = Item{
        Value:      value,
        Expiration: time.Now().Add(duration).UnixNano(),
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    item, found := c.items[key]
    if !found {
        return nil, false
    }

    if time.Now().UnixNano() > item.Expiration {
        return nil, false
    }

    return item.Value, true
}

func (c *Cache) startCleanupTimer() {
    ticker := time.NewTicker(12 * time.Hour)
    for range ticker.C {
        c.mu.Lock()
        for key, item := range c.items {
            if time.Now().UnixNano() > item.Expiration {
                delete(c.items, key)
            }
        }
        c.mu.Unlock()
    }
}