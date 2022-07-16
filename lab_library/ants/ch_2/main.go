package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/panjf2000/ants"
)

var sum int32

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

//例子一 使用普通的pool
func main() {
	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	p, _ := ants.NewPool(100)
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			demoFunc()
			wg.Done()
		})
	}
	wg.Wait()
	log.Printf("pool, capacity:%d", p.Cap())
	log.Printf("pool, running workers number:%d", p.Running())
	log.Printf("pool, free workers number:%d", p.Free())
}
