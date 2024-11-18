package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	map_cache map[string]cacheEntry
	mux       *sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		map_cache: make(map[string]cacheEntry),
		mux:       &sync.Mutex{}}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.map_cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cache_entry, exist := c.map_cache[key]
	return cache_entry.val, exist
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	previousTime := time.Now().UTC().Add(-interval)

	for k, v := range c.map_cache {
		if v.createdAt.Before(previousTime) {
			delete(c.map_cache, k)
		}
	}
}
