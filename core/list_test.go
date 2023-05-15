package core

import (
	"testing"
)

func TestList(t *testing.T) {
	START, END := 10, 20
	// init
	var list List
	list.Init()
	// add
	for i := START; i <= END; i++ {
		list.Add(CacheElem{
			key:        i,
			value:      i,
			expireTime: 0,
		})
	}
	// find
	for i := START; i <= END; i++ {
		_, code := list.Find(i)
		if code != HIT {
			t.Fatal("Test list fail.")
		}
	}
	// delete
	for i := START; i < (START+END)/2; i++ {
		list.Delete(i)
	}
	// find after delete
	for i := START; i <= END; i++ {
		_, code := list.Find(i)
		if i < (START+END)/2 && code == HIT {
			t.Fatal("Test list fail.")
		} else if i >= (START+END)/2 && code == MISS {
			t.Fatal("Test list fail.")
		}
	}
}
