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


