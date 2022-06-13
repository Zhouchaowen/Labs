package main

import (
	"fmt"
	"github.com/VictoriaMetrics/fastcache"
)

// Set/Get/Has/HasGet/Del 数据
// hash冲突覆盖
func main() {
	cacheSize := 1024
	cache := fastcache.New(cacheSize)

	cache.Set([]byte("key"), []byte("value"))
	cache.Set([]byte("key"), []byte("value_repeat"))

	v := cache.Get(nil, []byte("key"))
	fmt.Printf("Value:%s\n", v)

	exist := cache.Has([]byte("key"))
	fmt.Printf("Exist:%t\n", exist)

	v, exist = cache.HasGet(nil, []byte("key"))
	fmt.Printf("Value:%s, Exist:%t\n", v, exist)

	cache.Del([]byte("key"))

	v = cache.Get(nil, []byte("key"))
	fmt.Printf("Value:%s\n", v)
}
