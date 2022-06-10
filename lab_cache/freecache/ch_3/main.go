package main

import (
	"fmt"
	"github.com/coocood/freecache"
)

// GetFn
// 通过Key获取不到Value并对Value值做一些自定义操作
func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	key := []byte("key")
	val := []byte("value")
	expire := 60 // expire in 60 seconds

	cache.Set(key, val, expire)

	cache.GetFn(key, func(val []byte) error {
		fmt.Println("Value: ", string(val))
		return nil
	})

	fmt.Println("entry count ", cache.EntryCount())
}
