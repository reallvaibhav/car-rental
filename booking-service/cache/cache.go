package cache

import (
	"sync"
	"time"
)

type cacheItem struct {
	value      interface{}
	expiration int64 // Unix timestamp in nanoseconds
}

type Cache struct {
	data map[string]*cacheItem
	mu   sync.RWMutex
}

func NewCache() *Cache {
	c := &Cache{
		data: make(map[string]*cacheItem),
	}
	go c.startCleanupTimer()
	return c
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var exp int64
	if ttl > 0 {
		exp = time.Now().Add(ttl).UnixNano()
	} else {
		exp = 0 // 0 means no expiration
	}
	c.data[key] = &cacheItem{
		value:      value,
		expiration: exp,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.data[key]
	c.mu.RUnlock()
	if !found {
		return nil, false
	}
	if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
		// expired, delete and return not found
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
		return nil, false
	}
	return item.value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]*cacheItem)
}

func (c *Cache) startCleanupTimer() {
	ticker := time.NewTicker(1 * time.Hour)
	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.data {
			if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
