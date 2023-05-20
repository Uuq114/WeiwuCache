package core

import (
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

type Cache struct {
	listObject            *List
	hashObject            *map[string]CacheElem
	storageType           StorageType
	mutex                 sync.RWMutex
	defaultExpireDuration int64 // in second

	// todo: callback function when evict elem
}

func (cache *Cache) Init() {
	log.Println("init cache...")
	// load config
	cache.defaultExpireDuration = int64(viper.GetInt("DefaultExpireDuration"))
	// init
	if cache.storageType == LIST {
		cache.listObject.Init()
	} else if cache.storageType == DICT {
		//todo
	} else {
		log.Fatalln("Cache storage type not supported.")
	}
	//cache.elems = make(map[string]CacheElem)
}

func (cache *Cache) SetWithDefaultExpiration(key interface{}, value interface{}) bool {
	elem := NewElem(key, value, time.Now().Unix()+cache.defaultExpireDuration)
	//elem := CacheElem{Object: value, ExpireTime: time.Now().Unix() + cache.defaultExpireDuration}
	return cache.set(elem)
}

func (cache *Cache) SetWithExpiration(key interface{}, value interface{}, expiration int64) bool {
	elem := NewElem(key, value, expiration)
	//elem := CacheElem{Object: value, ExpireTime: time.Now().Unix() + expiration}
	return cache.set(elem)
}

func (cache *Cache) Get(key interface{}) interface{} {
	resp := cache.get(key)
	if resp.Ok() {
		return resp.Content()
	} else {
		return "miss or stale"
	}
}

func (cache *Cache) set(elem CacheElem) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if cache.storageType == LIST {
		cache.listObject.Add(elem)
	} else if cache.storageType == DICT {
		// hash set elem
	} else {
		log.Fatalln("Cache storage type not supported.")
	}
	// todo: add map size check, return false if size exceeds limit
	return true
}

func (cache *Cache) get(key interface{}) Response {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	var found interface{}
	var code RespCode
	if cache.storageType == LIST {
		found, code = cache.listObject.Find(key)
	} else if cache.storageType == DICT {
		// find in hash
	} else {
		log.Fatalln("Cache storage type not supported.")
	}

	switch code {
	case MISS:
		return Response{code: MISS, result: "nil"}
	case Stale:
		// todo: delete stale key from map
		return Response{code: Stale, result: found}
	case HIT:
		return Response{code: HIT, result: found}
	default:
		return Response{code: Error, result: nil}
	}
}
