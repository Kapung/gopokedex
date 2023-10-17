package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]entry
	mux   *sync.Mutex
}

type entry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]entry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = entry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reap() {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, value := range c.cache {
		if time.Since(value.createdAt) > (time.Minute * 5) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap()
	}
}
