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

// Defines a new cache type/initialization
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	// initialize go routine for cleanup of the cache
	go c.reapLoop()
	return c
}

// Cache method to write to cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = entry
}

// Cache method for Read
func (c *Cache) Get(key string) ([]byte, bool) {
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
	// Ticker with interval from cache
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	// Upon every interval of time, we clear the items in cache if they persist too long
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
