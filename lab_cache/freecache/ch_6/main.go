package main

import (
	"fmt"
	"github.com/coocood/freecache"
)

// Touch
// 验证通过key刷新缓存有效时间,被刷新的key必须还在缓存期类
func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	key1 := []byte("abcd")
	val1 := []byte("efgh")
	key2 := []byte("ijkl")
	val2 := []byte("mnop")
	err := cache.Set(key1, val1, 1)
	if err != nil {
		fmt.Printf("err should be nil %s\n", err.Error())
	}
	err = cache.Set(key2, val2, 1)
	if err != nil {
		fmt.Printf("err should be nil %s\n", err.Error())
	}

	err = cache.Touch(key1, 20)
	if err != nil {
		fmt.Printf("err should be nil %s\n", err.Error())
	}

	ttl, err := cache.TTL(key1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ttl value: %d\n", ttl)
	}

	fmt.Println("entry count ", cache.EntryCount())
}
