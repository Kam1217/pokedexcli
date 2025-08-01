package pokecashe

import (
	"sync"
	"time"
)

type cacheEntry struct {
    createdAt time.Time
    val       []byte	
}

type Cache struct {
	data map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache{
	// go cache.reapLoop()
	return &Cache{
		data: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cEntry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.data[key] = cEntry
}



