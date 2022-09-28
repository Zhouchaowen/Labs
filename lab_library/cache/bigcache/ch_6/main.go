// 测试待状态的删除回调 OnRemoveWithReason
package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func OnRemoveWithReason1() {
	onRemoveWR := func(key string, entry []byte, reason bigcache.RemoveReason) { // RemoveReason
		fmt.Printf("[callback] key:%s,vaule:%s removed. reason：%d\n", key, entry, reason)
	}
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             1, // 为了演示set时 剔除过期清除
		LifeWindow:         time.Second,
		OnRemoveWithReason: onRemoveWR,
	})

	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	err := cache.Set("key2", []byte("value2"))
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(value))
	}
}

func main() {
	OnRemoveWithReason1()
}
