package core

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	START, END := 10, 20
	// init
	var list List
	list.Init()
	fmt.Println("list init...")
	// add
	for i := START; i <= END; i++ {
		list.Add(CacheElem{
			key:        i,
			value:      i,
			expireTime: 0,
		})
	}
	fmt.Printf("list add elems, range: %d --> %d\n", START, END)
	fmt.Printf("list length: %d\n", list.Length)
	// find
	for i := START; i <= END; i++ {
		_, code := list.Find(i)
		if code != HIT {
			t.Fatal("Test list fail.")
		}
	}
	fmt.Printf("list find elems, range: %d --> %d\n", START, END)
	// delete
	for i := START; i < (START+END)/2; i++ {
		list.Delete(i)
	}
	fmt.Printf("list delete elems, range: %d --> %d\n", START, (START+END)/2-1)
	// find after delete
	for i := START; i <= END; i++ {
		_, code := list.Find(i)
		if i < (START+END)/2 && code == HIT {
			t.Fatalf("Test list fail, i: %d\n", i)
		} else if i >= (START+END)/2 && code == MISS {
			t.Fatalf("Test list fail, i: %d\n", i)
		}
	}
	fmt.Printf("list find elems, range: %d --> %d\n", START, END)
}
