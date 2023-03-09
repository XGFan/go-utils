package utils

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTTlCache(t *testing.T) {
	t.Run("expire after ttl", func(t *testing.T) {
		cache := NewTTlCache[string, string](time.Second * 1)
		cache.Set("hello", "world", 0)
		cache.Set("expire_after_2s", "value2", time.Second*2)
		get := cache.Get("hello")
		if get != "world" {
			t.Errorf("should get [world], but got %s", get)
		}
		time.Sleep(time.Second)
		get = cache.Get("hello")
		if get != "" {
			t.Errorf("should not get any value, but got %s", get)
		}
		get = cache.Get("expire_after_2s")
		if get != "value2" {
			t.Errorf("should get [value2], but got %s", get)
		}
		time.Sleep(time.Second)
		get = cache.Get("expire_after_2s")
		if get != "" {
			t.Errorf("should not get any value, but got %s", get)
		}
	})
}

func TestTTlCacheFilter(t *testing.T) {
	t.Run("expire after ttl", func(t *testing.T) {
		var counter int32 = 0
		cache := NewTTlCache[int, string](time.Hour * 1)
		mutex := sync.RWMutex{}
		mutex.Lock()
		for i := 0; i < 16; i++ {
			go func() {
				j := 1
				mutex.RLock()
				for {
					if cache.Filter(j % 100) {
						atomic.AddInt32(&counter, 1)
					}
					j++
				}
			}()
		}
		mutex.Unlock()
		time.Sleep(time.Second)
		if counter != 100 {
			t.Errorf("should get 100, but got %d", counter)
		}
	})
}
