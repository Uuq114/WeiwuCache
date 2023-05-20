package core

import (
	"hash/fnv"
	"log"
	"math"
	"math/rand"
)

type Entry struct {
	Key   interface{}
	Value interface{}
	next  *Entry
}

type Hashtable struct {
	table []*Entry
	size  int64
	mask  int64
	used  int64
}

type DictType struct {
	HashFunc func(key interface{}) int64
}

type Dict struct {
	DictType
	hashtable   [2]*Hashtable
	rehashIndex int64
}

func CreateDict(dictType DictType) *Dict {
	return &Dict{
		DictType: dictType,
		//hashtable:   [2]*Hashtable{nil, nil},
		rehashIndex: -1,
	}
}

func (dict *Dict) rehashRepeat() {
	dict.rehash(DICT_REHASH_REPEAT)
}

func (dict *Dict) rehash(recur int64) {
	for recur > 0 {
		recur -= 1
		if dict.hashtable[0].used == 0 {
			dict.hashtable[0], dict.hashtable[1] = dict.hashtable[1], nil
			dict.rehashIndex = -1
			return
		}
		// find first available list
		for dict.hashtable[dict.rehashIndex] == nil {
			dict.rehashIndex += 1
		}
		// migrate whole list
		entry := dict.hashtable[0].table[dict.rehashIndex]
		for entry != nil {
			entryNextCopy := entry.next
			pos := dict.HashFunc(entry.Key) & dict.hashtable[1].mask
			entry.next, dict.hashtable[1].table[pos] = dict.hashtable[1].table[pos], entry
			dict.hashtable[0].used -= 1
			dict.hashtable[1].used += 1
			entry = entryNextCopy
		}
	}
}

func (dict *Dict) isRehashing() bool {
	return dict.rehashIndex != -1
}

func (dict *Dict) expandIfNeeded() error {
	if dict.isRehashing() {
		return nil
	}
	if dict.hashtable[0] == nil {
		return dict.expand(HASHTABLE_INIT_SIZE)
	}
	if float64(dict.hashtable[0].used)/float64(dict.hashtable[0].size) > HASHTABLE_FULL_RATIO {
		return dict.expand(dict.hashtable[0].size * HASHTABLE_GROW_RATIO)
	}
	return nil
}

func (dict *Dict) expand(oldSize int64) error {
	newSize := getNewSize(oldSize)
	log.Printf("[INFO] expand, new size: %d\n", newSize)
	if dict.isRehashing() || (dict.hashtable[0] != nil && dict.hashtable[0].used > newSize) {
		log.Printf("[ERROR] expand hashtable fail\n")
		return EXPAND_HASHTABLE_ERR
	}
	hashtable := Hashtable{
		table: make([]*Entry, newSize),
		size:  newSize,
		mask:  newSize - 1,
		used:  0,
	}
	// if init dict
	if dict.hashtable[0] == nil {
		dict.hashtable[0] = &hashtable
		return nil
	}
	dict.hashtable[1] = &hashtable
	dict.rehashIndex = 0
	return nil
}

func getNewSize(oldSize int64) int64 {
	for size := HASHTABLE_INIT_SIZE; size < math.MaxInt64; size <<= 1 {
		if size > oldSize {
			return size
		}
	}
	return -1
}

/* CRUD */

func (dict *Dict) Set(key, value interface{}) error {
	if dict.isRehashing() {
		dict.rehashRepeat()
	}

	index := dict.getKeyPosition(key)
	if index != -1 {
		dict.addEntry(key, value)
	} else {
		entry := dict.find(key)
		entry.Value = value
	}
	log.Printf("[INFO] set entry, key: %v, value: %v\n", key, value)
	return nil
}

func (dict *Dict) Get(key interface{}) (interface{}, bool) {
	entry := dict.find(key)
	if entry == nil {
		return nil, false
	} else {
		return entry.Value, true
	}
}

func (dict *Dict) RandomGet() (interface{}, bool) {
	if dict.hashtable[0] == nil {
		return nil, false
	}
	hashtableIndex := 0
	if dict.isRehashing() {
		dict.rehashRepeat()
		// simply randomGet from bigger hashtable
		if dict.hashtable[1] != nil && dict.hashtable[1].used > dict.hashtable[0].used {
			hashtableIndex = 1
		}
	}
	hashtable := dict.hashtable[hashtableIndex]
	hashtableSize := hashtable.size
	var index int64
	for {
		index = rand.Int63n(hashtableSize)
		if hashtable.table[index] != nil {
			break
		}
	}
	var listLength int64
	for e := hashtable.table[index]; e != nil; e = e.next {
		listLength += 1
	}
	randomOffset := rand.Int63n(listLength)
	randomEntry := hashtable.table[index]
	for i := int64(0); i < randomOffset; i++ {
		randomEntry = randomEntry.next
	}
	return randomEntry.Value, true
}

func (dict *Dict) Delete(key interface{}) bool {
	if dict.hashtable[0] == nil {
		log.Printf("[INFO] delete: no such key, key: %v\n", key)
		return false
	}
	if dict.isRehashing() {
		dict.rehashRepeat()
	}
	hash := dict.HashFunc(key)
	for i := 0; i <= 1; i++ {
		index := hash & dict.hashtable[i].mask
		entry := dict.hashtable[i].table[index]
		var prev *Entry
		for entry != nil {
			//entryNextCopy := entry.next
			if entry.Key == key {
				if prev == nil {
					dict.hashtable[i].table[index] = nil
				} else {
					prev.next, entry.next = entry.next, nil
				}
				log.Printf("[INFO] delete: delete key, key: %v\n", key)
				return true
			}
			prev = entry
			entry = entry.next
		}
		if !dict.isRehashing() {
			break
		}
	}
	log.Printf("[INFO] delete: no such key, key: %v\n", key)
	return false
}

func (dict *Dict) addEntry(key, value interface{}) {
	if dict.isRehashing() {
		dict.rehashRepeat()
	}
	hash := dict.HashFunc(key)
	var hashtable *Hashtable
	if dict.isRehashing() {
		hashtable = dict.hashtable[1]
	} else {
		hashtable = dict.hashtable[0]
	}
	entry := &Entry{
		Key:   key,
		Value: value,
		next:  nil,
	}
	index := hash & hashtable.mask
	if hashtable.table[index] != nil {
		entry.next = hashtable.table[index]
	}
	hashtable.table[index] = entry
	hashtable.used += 1
}

func (dict *Dict) find(key interface{}) *Entry {
	if dict.hashtable[0] == nil {
		return nil
	}
	if dict.isRehashing() {
		dict.rehashRepeat()
	}
	hash := dict.HashFunc(key)
	var index int64
	for i := 0; i <= 1; i++ {
		index = hash & dict.hashtable[i].mask
		entry := dict.hashtable[i].table[index]
		for entry != nil {
			if entry.Key == key {
				return entry
			}
			entry = entry.next
		}
		if !dict.isRehashing() {
			break
		}
	}
	return nil
}

// get Hash(key) position
// return -1 if key already exists or expand() throws err
func (dict *Dict) getKeyPosition(key interface{}) int64 {
	if err := dict.expandIfNeeded(); err != nil {
		return -1
	}
	hash := dict.HashFunc(key)
	var index int64
	for i := 0; i <= 1; i++ {
		index = hash & dict.hashtable[i].mask
		entry := dict.hashtable[i].table[index]
		for entry != nil {
			if entry.Key == key {
				return -1
			}
			entry = entry.next
		}
		if !dict.isRehashing() {
			break
		}
	}
	return index
}

/* Hash function */
func hashCode(key interface{}) int64 {
	if key == nil {
		return 0
	} else {
		hash := fnv.New64()
		_, err := hash.Write([]byte(key.(string)))
		if err != nil {
			return 0
		}
		return int64(hash.Sum64())
	}
}
