package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				fmt.Println("clean up")
				quit <- "See you!" // main函数不消费会被阻塞
				fmt.Println("send See you ok")
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c
}

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := 3; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye"
	time.Sleep(3 * time.Second)
	fmt.Println("Joe say:", <-quit)
	time.Sleep(1 * time.Second)
}
