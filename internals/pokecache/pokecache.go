package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, exists := c.cache[key]

	return entry.val, exists
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mutex.Lock()

		for k, v := range c.cache {
			if time.Since(v.createdAt) >= interval {
				delete(c.cache, k)
			}
		}

		c.mutex.Unlock()
	}
}
