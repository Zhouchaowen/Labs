package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

//例子一 使用普通的pool
func main() {
	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	p, _ := ants.NewPool(100, ants.WithNonblocking(false))
	defer p.Release()

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		err := p.Submit(func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Hello World run %d wait %d \n", p.Running(), p.Waiting())
			wg.Done()
		})
		if err != nil {
			fmt.Println("err ", err.Error())
			wg.Done()
		}
	}
	wg.Wait()

	log.Printf("pool, capacity:%d", p.Cap())
	log.Printf("pool, running workers number:%d", p.Running())
	log.Printf("pool, free workers number:%d", p.Free())
}
