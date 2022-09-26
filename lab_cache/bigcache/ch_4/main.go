// CleanWindow 验证设置条目定期清理时间
package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func CleanWindow1() {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:      4,
		LifeWindow:  4 * time.Second,
		CleanWindow: time.Second,
	})
	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println("CleanWindow1: ", err.Error())
	} else {
		fmt.Println("CleanWindow1: ", string(value))
	}
}

func CleanWindow2() {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:      4,
		LifeWindow:  2 * time.Second,
		CleanWindow: time.Second,
	})
	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println("CleanWindow2: ", err.Error())
	} else {
		fmt.Println("CleanWindow2: ", string(value))
	}
}

func main() {
	CleanWindow1()
	CleanWindow2()
}
