package main

import (
	"context"
	"fmt"
	"time"
)

// the boring function return a channel to communicate with it.
func boring(ctx context.Context, msg string) <-chan string { // <-chan string means receives-only channel of string.
	c := make(chan string)
	go func() { // we launch goroutine inside a function.
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println(msg, " exit!")
				return
			case c <- fmt.Sprintf("%s %d", msg, i):
				fmt.Printf("send %s %d\n", msg, i)
				//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond) // 必要参数，否则无法正常关闭
			}
		}
	}()
	return c // return a channel to caller.
}

func fanIn(ctx context.Context, cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs { // spawn channel based on the number of input channel
		go func(cv <-chan string) { // cv is a channel value
			for {
				select {
				case <-ctx.Done():
					fmt.Println("fanIn exit!")
					return
				case c <- <-cv:

				}
			}
		}(ci) // send each channel to

	}
	return c
}

func main() {
	// 如果不调用cancel()会造成数据丢失
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel() // 进程安全退出
		fmt.Println("exec cancel()")
		time.Sleep(2 * time.Second)
	}()

	// merge 2 channels into 1 channel
	c := fanIn(ctx, boring(ctx, "Joe"), boring(ctx, "Ahn"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c) // now we can read from 1 channel
	}

	fmt.Println("You're both boring. I'm leaving")
}
