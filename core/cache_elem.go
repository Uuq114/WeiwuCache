package core

import "time"

type CacheElem struct {
	key        interface{}
	value      interface{}
	expireTime int64
}

func (elem *CacheElem) IsExpired() bool {
	if elem.expireTime == 0 {
		return false
	}
	return time.Now().Unix() > elem.expireTime
}

func NewDummyElem() CacheElem {
	return CacheElem{
		key:        "#",
		value:      "#",
		expireTime: 0,
	}
}

func NewElem(key interface{}, value interface{}, expiration int64) CacheElem {
	return CacheElem{
		key:        key,
		value:      value,
		expireTime: expiration,
	}
}
func (elem *CacheElem) Key() interface{} {
	return elem.key
}

func (elem *CacheElem) Value() interface{} {
	return elem.value
}
