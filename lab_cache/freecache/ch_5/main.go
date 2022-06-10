package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"time"
)

// TTL
// 验证缓存有效时间
func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	key := []byte("key")
	val := []byte("value")
	expire := 5 // expire in 60 seconds

	cache.Set(key, val, expire)

	time.Sleep(3 * time.Second)

	ttl, err := cache.TTL(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ttl value: %d\n", ttl)
	}

	fmt.Println("entry count ", cache.EntryCount())
}
