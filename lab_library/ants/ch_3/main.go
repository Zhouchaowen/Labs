package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"sync"
	"sync/atomic"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func main() {
	defer ants.Release()

	runTimes := 1000
	var wg sync.WaitGroup
	// Use the pool with a method,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()

	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	log.Printf("pool, capacity:%d", p.Cap())
	log.Printf("pool, running workers number:%d", p.Running())
	log.Printf("pool, free workers number:%d", p.Free())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	if sum != 499500 {
		panic("the final result is wrong!!!")
	}
}
