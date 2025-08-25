package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheEntries: map[string]cacheEntry{},
		mu:           &sync.Mutex{},
	}

	c.reapLoop(interval)

	return c
}
