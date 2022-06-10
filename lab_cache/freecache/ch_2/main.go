package main

import (
	"fmt"
	"github.com/coocood/freecache"
)

// GetOrSet
// 通过Key获取不到Value就将Value添加到缓存，获取的到就返回

// SetAndGet
// 通过Key将Value添加到缓存，并返回老的Value
func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	key := []byte("key")
	val := []byte("value")
	expire := 60 // expire in 60 seconds

	r, err := cache.GetOrSet(key, val, expire)
	if err != nil || r != nil {
		fmt.Printf("Expected to have val=%v and err != nil, got: value=%v, err=%v\n", string(val), string(r), err)
	}

	got, err := cache.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("get value: %s\n", got)
	}

	fmt.Println("=========================")
	// call twice for the same key
	valRepeat := []byte("xxxx")
	r, _, err = cache.SetAndGet(key, valRepeat, expire)
	if err != nil || string(r) != "value" {
		fmt.Printf("Expected to get old record, got: value=%v, err=%v\n", string(r), err)
	} else {
		fmt.Println("SetAndGet Value: ", string(r))
	}

	got, err = cache.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("get value: %s\n", got)
	}

	fmt.Println("entry count ", cache.EntryCount())
}
