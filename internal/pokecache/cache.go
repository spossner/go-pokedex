package pokecache

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	lock    sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		lock:    sync.RWMutex{},
	}
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			c.reapLoop(interval)
		}
	}()
	return c
}
func (c *Cache) Add(key string, data []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(duration time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > duration {
			delete(c.entries, key)
		}
	}
}

func (c *Cache) GetUrl(url string) ([]byte, error) {
	var data []byte
	data, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error getting next page: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response: %w", err)
		}
		c.Add(url, data)
	}
	return data, nil
}
