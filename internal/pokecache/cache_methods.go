package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cacheEntries[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for t := range ticker.C {
			c.mu.Lock()
			for key, val := range c.cacheEntries {
				lastTick := t.Add(-interval)
				if val.createdAt.Before(lastTick) {
					delete(c.cacheEntries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
