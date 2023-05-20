package core

import "errors"

type RespCode int

const (
	MISS RespCode = iota
	HIT
	Stale
	Error
)

type StorageType int

const (
	LIST StorageType = iota
	DICT
)

var (
	EXPAND_HASHTABLE_ERR = errors.New("expand hashtable error")
)

const (
	HASHTABLE_INIT_SIZE  int64   = 8
	HASHTABLE_FULL_RATIO float64 = 1.5
	HASHTABLE_GROW_RATIO int64   = 2
	DICT_REHASH_REPEAT   int64   = 1
)
