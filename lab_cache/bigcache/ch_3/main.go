package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

// 测试过期删除 LifeWindow
func main() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Second))
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
