package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

func (item CacheItem) IsExpired() bool {
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	items map[string]CacheItem
	mutex sync.RWMutex
}

func New() *Cache {
	cache := &Cache{
		items: make(map[string]CacheItem),
	}

	// Start cleanup goroutine
	go cache.cleanup()

	return cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expiration := time.Now().Add(duration).UnixNano()
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.IsExpired() {
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}

func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items = make(map[string]CacheItem)
}

func (c *Cache) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.deleteExpired()
		}
	}
}

func (c *Cache) deleteExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, item := range c.items {
		if item.IsExpired() {
			delete(c.items, key)
		}
	}
}
