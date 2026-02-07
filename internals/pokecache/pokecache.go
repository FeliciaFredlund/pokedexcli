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
	mutex *sync.Mutex
}

func (c Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mutex.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	entry, exists := c.cache[key]
	c.mutex.Unlock()

	return entry.val, exists
}

func (c Cache) reapLoop(interval time.Duration) {

}

func NewCache(interval time.Duration) Cache {

}
