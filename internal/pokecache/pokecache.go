package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	mu    sync.Mutex
	cache map[string]cacheEntry
	ttl   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewPokeCache(ttl time.Duration) *PokeCache {
	pc := &PokeCache{
		cache: make(map[string]cacheEntry),
		ttl:   ttl,
	}

	go pc.reapLoop(ttl)
	return pc
}

func (pc *PokeCache) Get(key string) ([]byte, bool) {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	entry, ok := pc.cache[key]
	if !ok {
		return nil, false
	}

	if time.Since(entry.createdAt) > pc.ttl {
		delete(pc.cache, key)
		return nil, false
	}

	return entry.val, true
}

func (pc *PokeCache) Add(key string, data []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	newCacheEntry := cacheEntry{
		val:       data,
		createdAt: time.Now(),
	}

	pc.cache[key] = newCacheEntry
}

func (pc *PokeCache) reapLoop(i time.Duration) {
	ticker := time.NewTicker(i)
	for range ticker.C {
		go pc.reap(i)
	}
}

func (pc *PokeCache) reap(i time.Duration) {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	for key, entry := range pc.cache {
		if (time.Since(entry.createdAt)) > i {
			delete(pc.cache, key)
		}
	}
}
