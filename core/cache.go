package core

import (
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

type Cache struct {
	elems                 map[string]CacheElem
	mutex                 sync.RWMutex
	defauluExpireDuration int64 // in second

	// todo: callback function when evict elem
}

func (cache *Cache) Init() {
	log.Println("init cache...")
	cache.elems = make(map[string]CacheElem)
	cache.defauluExpireDuration = int64(viper.GetInt("DefaultExpireDuration"))
}

func (cache *Cache) SetWithDefaultExpiration(key string, value interface{}) bool {
	elem := CacheElem{Object: value, ExpireTime: time.Now().Unix() + cache.defauluExpireDuration}
	return cache.set(key, elem)
}

func (cache *Cache) SetWithExpiration(key string, value interface{}, expiration int64) bool {
	elem := CacheElem{Object: value, ExpireTime: time.Now().Unix() + expiration}
	return cache.set(key, elem)
}

func (cache *Cache) Get(key string) interface{} {
	resp := cache.get(key)
	if resp.Ok() {
		return resp.Content()
	} else {
		return "miss or stale"
	}
}

func (cache *Cache) set(key string, elem CacheElem) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.elems[key] = elem
	// todo: add map size check, return false if size exceeds limit
	return true
}

func (cache *Cache) get(key string) Response {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	elem, ok := cache.elems[key]
	if !ok {
		return Response{code: MISS, result: "nil"}
	} else {
		if elem.IsExpired() {
			// todo: delete stale key from map
			return Response{code: Stale, result: elem.Value()}
		} else {
			return Response{code: HIT, result: elem.Value()}
		}
	}
}
