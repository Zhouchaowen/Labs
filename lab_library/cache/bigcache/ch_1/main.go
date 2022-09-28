// Set/Get/Del 数据
package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func main() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value"))

	cache.Set("my-unique-key", []byte("value1"))

	entry, err := cache.Get("my-unique-key")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Value:%s\n", string(entry))
	}

	cache.Delete("my-unique-key")
	entry, err = cache.Get("my-unique-key")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Value:%s\n", string(entry))
	}
}
