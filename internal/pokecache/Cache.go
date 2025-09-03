package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntries map[string]CacheEntry
	mu           *sync.Mutex
}

func (ch *Cache) Add(key string, val []byte) {
	cache := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	ch.mu.Lock()
	ch.cacheEntries[key] = cache
	ch.mu.Unlock()
}

func (ch *Cache) Get(key string) ([]byte, bool) {
	ch.mu.Lock()
	value, ok := ch.cacheEntries[key]
	ch.mu.Unlock()
	if ok {
		return value.val, true
	} else {
		return nil, false
	}
}

func (ch *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		ch.mu.Lock()
		for key, cache := range ch.cacheEntries {
			timeSince := time.Since(cache.createdAt)
			if timeSince > interval {
				delete(ch.cacheEntries, key)
			}
		}
		ch.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cacheEntries := make(map[string]CacheEntry)
	var mtx sync.Mutex
	newCh := Cache{
		cacheEntries: cacheEntries,
		mu:           &mtx,
	}
	go newCh.reapLoop(interval)
	return &newCh
}
