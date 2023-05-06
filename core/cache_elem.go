package core

import "time"

type CacheElem struct {
	Object     interface{}
	ExpireTime int64
}

func (elem *CacheElem) IsExpired() bool {
	if elem.ExpireTime == 0 {
		return false
	}
	return time.Now().Unix() > elem.ExpireTime
}

func (elem *CacheElem) Value() interface{} {
	return elem.Object
}
