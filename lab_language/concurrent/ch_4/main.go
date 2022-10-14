// 通过channel退出goroutines
// https://blog.csdn.net/tianmaxingkong_/article/details/104229878
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		defer func() {
			fmt.Println(msg, "goroutines exit!")
			if err := recover(); err != nil {
				fmt.Println("err:", err)
			}
		}()

		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case str := <-quit:
				fmt.Println("recover:", str)
				quit <- msg + " See you main!" // main函数不消费会被阻塞
				fmt.Println(msg, "send See you ok")
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e4)) * time.Millisecond)
		}
	}()
	return c
}

func deadlock() {
	quit := make(chan string)
	jc := boring("Joe", quit)
	ac := boring("Ale", quit)
	for i := 3; i >= 0; i-- {
		fmt.Println(<-jc, <-ac)
	}
	quit <- "Bye" // 巨坑 会导致deadlock。
	time.Sleep(3 * time.Second)
	fmt.Println("say:", <-quit)
	fmt.Println("say:", <-quit) // 巨坑 会导致deadlock。
	time.Sleep(10 * time.Second)
}

func noDeadlock() {
	quit := make(chan string)
	jc := boring("Joe", quit)
	ac := boring("Ale", quit)
	for i := 3; i >= 0; i-- {
		fmt.Println(<-jc, <-ac)
	}

	close(quit) // 所有都会收到
	time.Sleep(3 * time.Second)
	fmt.Println("say:", <-quit)
	fmt.Println("say:", <-quit)
	time.Sleep(10 * time.Second)
}

// 多个go协程监听一个chan时，同一时刻 向chan写入数据时，只有一个go协程会获取到数据
// 多个go协程监听一个chan时，同一时刻 关闭chan，所有go协程会从chan获取到关闭通知
func main() {
	//deadlock()
	noDeadlock()
}
