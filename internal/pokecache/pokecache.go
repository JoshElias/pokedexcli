package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data         map[string]cacheEntry
	mu           sync.Mutex
	reapInterval time.Duration
	maxAge       time.Duration
	done         chan struct{}
	isReaping    bool
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		reapInterval: interval,
		maxAge:       7 * time.Second,
		done:         make(chan struct{}),
		isReaping:    false,
	}
	return &cache
}

func (c *Cache) addLocked(key string, val []byte) {
	wasEmpty := len(c.data) == 0

	if c.data == nil {
		c.data = make(map[string]cacheEntry)
	}

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	if wasEmpty {
		go c.reapLoop()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.addLocked(key, val)
}

func (c *Cache) getLocked(key string) ([]byte, bool) {
	if entry, ok := c.data[key]; !ok {
		return nil, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.getLocked(key)
}

func (c *Cache) deleteLocked(key string) bool {
	_, exists := c.data[key]
	delete(c.data, key)

	if len(c.data) == 0 {
		c.stopReaping()
	}

	return exists
}

func (c *Cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.deleteLocked(key)
}

func (c *Cache) reapLoop() {
	c.isReaping = true
	for {
		select {
		case <-c.done:
			c.isReaping = false
			return
		default:
			c.reap()
			time.Sleep(c.reapInterval)
		}
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.data {
		oldestAllowed := entry.createdAt.Add(c.maxAge)
		if oldestAllowed.Before(now) {
			c.deleteLocked(key)
		}
	}
}

func (c *Cache) stopReaping() {
	if c.isReaping {
		go func() { c.done <- struct{}{} }()
	}
}

func (c *Cache) IsReaping() bool {
	return c.isReaping
}

// func (c *Cache) Destroy() {
// 	fmt.Println("destroying!!")
//
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
//
// 	if !c.isReaping {
// 		return
// 	}
//
// 	select {
// 	case c.done <- struct{}{}:
// 		// Signal sent successfully
// 	default:
// 		// Channel already signaled
// 	}
// 	close(c.done)
// 	fmt.Println("Destroyed Cache")
// }
