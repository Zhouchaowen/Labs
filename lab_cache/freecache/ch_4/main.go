package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"time"
)

// Expire
// 验证缓存过期时间
func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	key := []byte("key")
	val := []byte("value")
	expire := 2 // expire in 60 seconds

	cache.Set(key, val, expire)

	time.Sleep(5 * time.Second)

	got, err := cache.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("get value: %s\n", got)
	}

	fmt.Println("entry count ", cache.EntryCount())
}
