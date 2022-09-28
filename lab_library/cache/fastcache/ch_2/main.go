package main

import (
	"fmt"
	"github.com/VictoriaMetrics/fastcache"
)

// Set/Get Big数据
// 测试大数据缓存插入和查询
func main() {
	cacheSize := 1024
	cache := fastcache.New(cacheSize)

	// Both key and value exceed 64Kb
	k := make([]byte, 90*1024)
	v := make([]byte, 100*1024)
	cache.Set(k, v)
	vv := cache.Get(nil, k)
	fmt.Printf("set > 64k len:%d\n", len(vv))

	// len(key) + len(value) > 64Kb
	k = make([]byte, 40*1024)
	v = make([]byte, 40*1024)
	cache.Set(k, v)
	vv = cache.Get(nil, k)
	fmt.Printf("set < 64k len:%d\n", len(vv))
}
