package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func main() {
	onRemove := func(key string, entry []byte) {
		fmt.Printf("key:%s,vaule:%s removed\n", key, entry)
	}

	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             4,
		CleanWindow:        time.Second,
		MaxEntriesInWindow: 1,
		MaxEntrySize:       256,
		OnRemove:           onRemove,
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
