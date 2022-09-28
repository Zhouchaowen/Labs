package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

// 测试不设置 LifeWindow
// 不设置LifeWindow的话会导致最先写入的数据，在load缓存数据的过程中就失效被剔除。
func main() {
	onRemove := func(key string, entry []byte) {
		fmt.Printf("[callback] key:%s,vaule:%s removed\n", key, entry)
	}

	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:   1,
		OnRemove: onRemove,
	})

	// when
	cache.Set("key", []byte("value"))
	time.Sleep(500 * time.Millisecond)
	cache.Set("key1", []byte("value1"))
	time.Sleep(500 * time.Millisecond)
	cache.Set("key2", []byte("value2"))
	time.Sleep(500 * time.Millisecond)
	cache.Set("key3", []byte("value3"))
	time.Sleep(3 * time.Second)
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(value))
	}
}
