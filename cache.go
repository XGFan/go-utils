package utils

import (
	"sync"
	"time"
)

type TTLCache[K any, V any] struct {
	ttl       time.Duration
	storage   sync.Map
	lock      sync.Mutex
	nextClean time.Time
}

type CacheItem[V any] struct {
	validBefore time.Time
	value       V
}

func (cache *TTLCache[K, V]) Get(key K) V {
	c, exist := cache.storage.Load(key)
	now := time.Now()
	if exist {
		cacheItem := c.(CacheItem[V])
		if cacheItem.validBefore.After(now) {
			return cacheItem.value
		} else {
			cache.storage.Delete(key)
		}
	}
	if cache.nextClean.Before(now) {
		if cache.lock.TryLock() {
			defer cache.lock.Unlock()
			cache.storage.Range(func(k, v any) bool {
				if v.(CacheItem[V]).validBefore.Before(now) {
					cache.storage.Delete(k)
				}
				return true
			})
			cache.nextClean = now.Add(cache.ttl)
		}
	}
	var result V
	return result
}

func (cache *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) {
	if ttl == 0 { //未设置ttl，使用默认ttl
		ttl = cache.ttl
	}
	if ttl > 0 {
		cache.storage.Store(
			key,
			CacheItem[V]{
				validBefore: time.Now().Add(ttl),
				value:       value,
			},
		)
	}
}

func (cache *TTLCache[K, V]) Filter(id K) bool {
	now := time.Now()
	actual, loaded := cache.storage.LoadOrStore(id, CacheItem[V]{
		validBefore: now.Add(cache.ttl),
	})
	if loaded {
		return actual.(CacheItem[V]).validBefore.Before(now)
	} else {
		return true
	}

}

func NewTTlCache[K comparable, V comparable](ttl time.Duration) *TTLCache[K, V] {
	dnsCache := &TTLCache[K, V]{
		ttl:       ttl,
		storage:   sync.Map{},
		lock:      sync.Mutex{},
		nextClean: time.Now().Add(ttl),
	}
	return dnsCache
}
