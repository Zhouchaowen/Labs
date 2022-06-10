package main

import (
	"fmt"
	"github.com/coocood/freecache"
)

// Set/Get/Del 数据
// hash冲突
func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	key := []byte("key")
	val := []byte("value")
	valRepeat := []byte("valueRepeat")
	expire := 60 // expire in 60 seconds

	cache.Set(key, val, expire)

	cache.Set(key, valRepeat, expire) // hash冲突

	got, err := cache.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("get value: %s\n", got)
	}

	affected := cache.Del(key)

	fmt.Println("deleted key ", affected)
	fmt.Println("entry count ", cache.EntryCount())
}
