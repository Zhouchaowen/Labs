// 原子操作
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	return make(map[string]string)
}

// 示例展示了如何使用 Value 定期更新程序配置
func UpdateConfig() {
	var config atomic.Value // holds current server configuration
	// Create initial config value and atomic into config.
	config.Store(loadConfig())
	go func() {
		// Reload config every 10 seconds
		// and update config value with the new version.
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()
	// Create worker goroutines that handle incoming requests
	// using the latest config value.
	for i := 0; i < 10; i++ {
		go func() {
			for {
				c := config.Load()
				fmt.Printf("%p\n", c)
				time.Sleep(5 * time.Second)
			}
		}()
	}
	select {}
}

// 如何使用写时复制习语维护可扩展的频繁读取但不经常更新的数据结构。
// 原子方式更新数据
func readMostly() {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex // used only by writers

	// read function can be used to read the data without further synchronization
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}

	// insert function can be used to update the data without further synchronization
	insert := func(key, val string) {
		mu.Lock() // synchronize with other potential writers
		defer mu.Unlock()
		m1 := m.Load().(Map) // load current value of the data structure
		m2 := make(Map)      // create a new value
		for k, v := range m1 {
			m2[k] = v // copy all data from the current object to the new one
		}
		m2[key] = val // do the update that we need
		m.Store(m2)   // atomically replace the current object with the new one
		// At this point all new readers start working with the new version.
		// The old version will be garbage collected once the existing readers
		// (if any) are done with it.
	}
	_, _ = read, insert
}

func atomicParamValue() {
	var countVal atomic.Value
	countVal.Store([]int{1, 3, 5, 7})
	anotherStore(countVal)
	fmt.Printf("The count value: %+v \n", countVal.Load())
}

func anotherStore(countVal atomic.Value) {
	countVal.Store([]int{2, 4, 6, 8})
}

func main() {
	//UpdateConfig()
	atomicParamValue()
}
