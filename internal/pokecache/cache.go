package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

// Create empty Cache and stores its pointer in config --> then start a background goroutine to remove(reap) older cache than interval time
// then return the config (same cache pointer - now only existing cache not older than interval time)
func NewCache(interval time.Duration) *Cache {
	config := &Cache{
		cache: make(map[string]cacheEntry),
	}
	go config.reapLoop(interval)

	return config
}

// Warning: Maps are not thread-safe in Go - use sync.Mutex to lock access to the map when operating with their method

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.cache {
			// should only delete entries that are OLDER than one interval, not all entries
			// Because " now.After(entry.createdAt) " is delete all entries after createdAt time

			// change to " now.Sub(entry.createdAt) >= interval "
			if now.Sub(entry.createdAt) >= interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
