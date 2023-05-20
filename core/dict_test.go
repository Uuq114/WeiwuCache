package core

import (
	"strconv"
	"testing"
)

func TestDict(t *testing.T) {
	// init
	dict := CreateDict(DictType{HashFunc: hashCode})
	// add
	START, END := 1, 20
	for i := START; i <= END; i++ {
		err := dict.Set(strconv.Itoa(i), i)
		if err != nil {
			return
		}
	}
	// find
	for i := START; i <= END; i++ {
		_, ok := dict.Get(strconv.Itoa(i))
		if !ok {
			t.Fatal("test dict find fail.")
		}
	}
	// random get
	for i := 0; i < (END-START+1)/2; i++ {
		value, ok := dict.RandomGet()
		if !ok || value.(int) < START || value.(int) > END {
			t.Fatal("test dict random get fail.")
		}
	}
	// delete
	for i := START; i < (START+END)/2; i++ {
		ok := dict.Delete(strconv.Itoa(i))
		if !ok {
			t.Fatal("test dict delete fail.")
		}
	}
	// find after delete
	for i := (START + END) / 2; i <= END; i++ {
		_, ok := dict.Get(strconv.Itoa(i))
		if !ok {
			t.Fatal("test dict find after delete fail.")
		}
	}
}
