package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

const interval = 500

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * interval * time.Millisecond)
		fmt.Println("cancelling context from 5 seconds")
		cancel()
	}()

	f(ctx) // 当协程执行cancel后f()将退出
}

func f(ctx context.Context) {
	ticker := time.NewTicker(interval * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			doSomething()
		case <-ctx.Done(): // cancel()调用后将接收到
			return
		}
	}
}

func doSomething() { log.Println("tick") }
