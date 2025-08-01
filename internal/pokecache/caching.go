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
	cach := &Cache{
		data: make(map[string]cacheEntry),
	}
	go cach.reapLoop(interval)
	return cach
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

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(-interval)
	defer ticker.Stop()

	for{
		<-ticker.C
		c.mu.Lock()

		for k, entry := range c.data{
			if entry.createdAt.Before(time.Now().Add(interval)){
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}

