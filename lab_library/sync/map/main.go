package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Demo struct {
	A string
	B *string
	C int
}

type Cache struct {
	Name   string
	Caches *sync.Map
}

func main() {
	cache := Cache{
		"zcw",
		new(sync.Map),
	}

	key := "123456789"
	key2 := "123456789"
	go func() {
		value1 := &Demo{
			C: 1,
		}
		cache.Caches.Store(key, value1)
	}()

	for i := 1; i < 5; i++ {
		runtime.GC()
		time.Sleep(5 * time.Second)
	}

	time.Sleep(1200 * time.Second)
	v, ok := cache.Caches.Load(key2)
	fmt.Printf("value:%+v, ok:%t", v, ok)
}
