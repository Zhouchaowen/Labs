package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func main() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value1"))

	cache.Append("my-unique-key", []byte("value2"))

	cache.Append("my-unique-key", []byte("value3"))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))
}
