package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"time"
)

// Iterator
// 验证 迭代遍历整个cache
func main() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	count := 10
	for i := 0; i < count; i++ {
		err := cache.Set([]byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("val%d", i)), 0)
		if err != nil {
			fmt.Printf("err should be nil %s\n", err.Error())
		}
	}
	// Set some value that expires to make sure expired entry is not returned.
	cache.Set([]byte("abc"), []byte("def"), 2)

	time.Sleep(5 * time.Second)

	it := cache.NewIterator()
	for i := 0; i < count; i++ {
		entry := it.Next()
		if entry == nil {
			fmt.Printf("entry is nil for %d\n", i)
		}
		fmt.Printf("entry key value not match %s %s\n", entry.Key, entry.Value)
	}

	fmt.Println("entry count ", cache.EntryCount())
}
