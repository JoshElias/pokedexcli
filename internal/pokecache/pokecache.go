package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data         map[string]cacheEntry
	mu           sync.Mutex
	reapInterval time.Duration
	reapAge      time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		reapInterval: interval,
		reapAge:      7 * time.Second,
	}
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data == nil {
		c.data = make(map[string]cacheEntry)
	}

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.data[key]; !ok {
		return nil, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) Delete(key string) bool {
	fmt.Println("deleting from cache: ", key)
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.data[key]
	delete(c.data, key)
	return exists
}

func (c *Cache) reapLoop() {
	timer := time.NewTimer(c.reapInterval)
	<-timer.C
	for key, entry := range c.data {
		age := entry.createdAt.Add(c.reapAge)
		if age.Before(time.Now()) {
			c.Delete(key)
		}
	}
}
