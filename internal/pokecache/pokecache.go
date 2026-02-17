package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    *sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	cache := Cache{cache: make(map[string]cacheEntry), mu: &sync.Mutex{}}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{val: val, createdAt: time.Now().UTC()}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cache[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	// setup a ticker
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for key := range c.cache {
			entry := c.cache[key]
			timePassed := time.Since(entry.createdAt)
			if timePassed > interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}

}
