// 读写锁 大量并发读，少量并发写
package main

import (
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	count uint64
}

// Incr 写保护
func (c *Counter) Incr()  {
	c.Lock()
	defer c.Unlock()
	c.count++
}

// Count 读保护
func (c *Counter) Count() uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.count
}

func main() {
	var counter Counter

	for i:=0;i<10;i++{
		go func() {
			counter.Count()
			time.Sleep(time.Millisecond)
		}()
	}

	for  {
		counter.Incr()
		time.Sleep(time.Second)
	}
}