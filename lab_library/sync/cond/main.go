// 等待唤醒
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var sharedRsc = false

func cond_1() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait()
		}
		fmt.Println("goroutine2", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	// this one writes changes to sharedRsc
	time.Sleep(2 * time.Second)
	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	c.Broadcast() // 唤醒协程
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()
	wg.Wait()
}

// 区别 cond_2 方法,改变
// 导致死锁
func cond_2() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		fmt.Println("goroutine2 wait")
		c.Wait()
		fmt.Println("goroutine2", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	// this one writes changes to sharedRsc
	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()
	wg.Wait()
}

func run(){
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0;i<10;i++{
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10))*time.Second)
			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()
			log.Printf("运动员#%d 已准备就绪\n",i)

			// 广播唤醒所有等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		 c.Wait()
		 log.Printf("裁判被唤醒\n")
	}
	c.L.Unlock()

	log.Printf("所有运动员都准备好了，准备开始.....\n")
}

func main() {
	//cond_1()
	//cond_2()
	run()
}

/*
https://ieevee.com/tech/2019/06/15/cond.html
https://stackoverflow.com/questions/36857167/how-to-correctly-use-sync-cond#
https://geektutu.com/post/hpg-sync-cond.html
*/
