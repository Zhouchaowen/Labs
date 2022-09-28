// LifeWindow 验证设置条目过期时间
package main

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func LifeWindow1() {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:      1, // 为了演示set时 剔除过期清除
		LifeWindow:  time.Second,
		CleanWindow: 0,
	})
	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println("LifeWindow1: ", err.Error())
	} else {
		fmt.Println("LifeWindow1: ", string(value))
	}
}

func LifeWindow2() {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:      1, // 为了演示set时 剔除过期清除
		LifeWindow:  time.Second,
		CleanWindow: 0,
	})
	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	cache.Set("key2", []byte("value2")) // set时将key的过期数据剔除
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println("LifeWindow2: ", err.Error())
	} else {
		fmt.Println("LifeWindow2: ", string(value))
	}
}

func LifeWindow3() {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:     1,
		LifeWindow: 4 * time.Second,
	})
	// when
	cache.Set("key", []byte("value"))
	<-time.After(3 * time.Second)
	cache.Set("key2", []byte("value2"))
	value, err := cache.Get("key")
	if err != nil {
		fmt.Println("LifeWindow3: ", err.Error())
	} else {
		fmt.Println("LifeWindow3: ", string(value))
	}
}

func main() {
	LifeWindow1()
	LifeWindow2()
	LifeWindow3()
}
