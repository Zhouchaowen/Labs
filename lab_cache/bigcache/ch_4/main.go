package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func main() {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             4,
		CleanWindow:        time.Second,
		MaxEntriesInWindow: 1,
		MaxEntrySize:       256,
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
