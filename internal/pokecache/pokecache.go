package pokecache

import (
	"sync"
	"time"


)

type Cache struct{
	big_cache map[string]cacheEntry
	mu 	*sync.Mutex
}

type cacheEntry struct{
	createdAt  time.Time
	val []byte
}


func NewCache(interval time.Duration) Cache {
	c := Cache{
		big_cache: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go c.reapLoop(interval)
    return c
}

func (C *Cache) Add(key string , val []byte){
	C.mu.Lock()
	defer C.mu.Unlock()
	C.big_cache[key] = cacheEntry{
			createdAt: time.Now(),
			val: val,
		}
	
}

func (C *Cache) Get(key string  )([]byte , bool){
	C.mu.Lock()
	defer C.mu.Unlock()
	entry, ok := C.big_cache[key]
	if !ok {
		return nil , false
	}else{
		return entry.val , true 
	} 

}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for t := range ticker.C {
		c.reap(t,interval)
    }
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
	threshold := now.Add(-interval)
    for key, entry := range c.big_cache {
		if entry.createdAt.Before(threshold) {
			delete(c.big_cache,key)
		}
    }
}