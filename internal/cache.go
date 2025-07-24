package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.RWMutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// constructor
func NewCache(interval time.Duration) *Cache {
	c := &Cache{

		entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

// reap loop
func (c *Cache) reapLoop(interval time.Duration) {
	//created a NewTicker
	//this ticker will "proc" every time the interval is passed
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		//wait for the next interval
		<-ticker.C
		c.mu.Lock()
		//delete all entries that are older than the interval
		for key, value := range c.entries {
			age := time.Since(value.createdAt)
			if age > interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

// add entry
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{time.Now(), val}
	c.entries[key] = entry
}

// get entry
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, ok := c.entries[key]; ok {
		return val.val, true
	}
	return nil, false
}
