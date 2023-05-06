package core

type RespCode int

const (
	MISS RespCode = iota
	HIT
	Stale
)
