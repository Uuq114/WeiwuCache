package main

import (
	"github.com/Uuq114/WeiwuCache/core"
	"log"
)

func main() {
	cache := core.Cache{}
	cache.Init()

	cache.SetWithDefaultExpiration("caocao", "weiwu")
	result := cache.Get("caocao")
	log.Println("result:", result)
}
