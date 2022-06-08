package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

// 测试定时清理 CleanWindow
func main() {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:      4,
		CleanWindow: time.Second,
	})
	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(value))
	}
}
