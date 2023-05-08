package core

import "time"

type CacheElem struct {
	object     interface{}
	expireTime int64
}

func (elem *CacheElem) IsExpired() bool {
	if elem.expireTime == 0 {
		return false
	}
	return time.Now().Unix() > elem.expireTime
}

func (elem *CacheElem) Value() interface{} {
	return elem.object
}

func NewDummyElem() CacheElem {
	return CacheElem{
		object:     "#",
		expireTime: 0,
	}
}

func NewElem(content interface{}, time int64) CacheElem {
	return CacheElem{
		object:     content,
		expireTime: time,
	}
}
