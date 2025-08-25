package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{}

	entry.createdAt = time.Now()
	entry.val = val

	c.mu.Lock()
	c.cacheEntries[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.cacheEntries[key]
	c.mu.Unlock()
	if !exists {
		return nil, false
	}

	return entry.val, true
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
