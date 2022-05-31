package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func main() {
	onRemoveExt := func(key string, entry []byte, reason bigcache.RemoveReason) {
		fmt.Printf("key:%s,vaule:%s removed. reason：%d\n", key, entry, reason)
	}
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             1,
		MaxEntriesInWindow: 1,
		MaxEntrySize:       3,
		OnRemoveWithReason: onRemoveExt,
	})

	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	err := cache.Set("key1", []byte("value11"))
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(value))
	}
}
