package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	// Creates a new cache
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	// Adds new item to cache
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = entry
}
func (c *Cache) Get(key string) ([]byte, bool) {
	// Receives data from the cache
	c.mu.Lock()
	defer c.mu.Unlock()

	// Checks if key exists and handles Key Error
	if _, ok := c.entries[key]; !ok {
		return nil, false
	}
	return c.entries[key].val, true
}

// Reap loop function cleans the cache if data is in cache for too long
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		c.mu.Lock()
		for k, e := range c.entries {
			if now.Sub(e.createdAt) > c.interval {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
