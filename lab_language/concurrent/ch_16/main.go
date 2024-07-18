package main

import (
	"fmt"
	"sync"
	"time"
)

// 测试close chan后是不是所有地方的chanel都会收到通知

func A(c chan string) {
	for {
		select {
		case <-c:
			fmt.Printf("return A\n")
			return
		}
	}
}

func B(c chan string) {
	for {
		select {
		case <-c:
			fmt.Printf("return B\n")
			return
		}
	}
}

func main() {
	c := make(chan string)
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		wg.Add(1)
		A(c)
	}()

	go func() {
		defer wg.Done()
		wg.Add(1)
		B(c)
	}()

	time.Sleep(5 * time.Second)
	close(c)

	wg.Wait()
	fmt.Printf("return\n")

	time.Sleep(5 * time.Second)
}
