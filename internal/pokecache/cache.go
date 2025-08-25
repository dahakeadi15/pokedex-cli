package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	interval     time.Duration
	mu           *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	mu := &sync.Mutex{}

	c := Cache{
		cacheEntries: map[string]cacheEntry{},
		interval:     interval,
		mu:           mu,
	}

	c.reapLoop()

	return c
}
