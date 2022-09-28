// 测试删除回调 onRemove
package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func onRemove1() {
	onRemove := func(key string, entry []byte) {
		fmt.Printf("[callback] key:%s,vaule:%s removed\n", key, entry)
	}

	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:      4,
		CleanWindow: time.Second,
		OnRemove:    onRemove,
	})

	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println("onRemove1: ", err.Error())
	} else {
		fmt.Println("onRemove1: ", string(value))
	}
}

func onRemove2() {
	onRemove := func(key string, entry []byte) {
		fmt.Printf("[callback] key:%s,vaule:%s removed\n", key, entry)
	}

	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:     1, // 为了演示set时 剔除过期清除
		LifeWindow: time.Second,
		OnRemove:   onRemove,
	})

	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	cache.Set("key2", []byte("value2"))
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println("onRemove2: ", err.Error())
	} else {
		fmt.Println("onRemove2: ", string(value))
	}
}

func main() {
	onRemove1()
	onRemove2()
}
