package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan string
}

func fanIn(ctx context.Context, inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("fanIn exit!")
					return
				default:
					c <- <-input
				}
			}
		}()
	}
	return c
}

// the boring function return a channel to communicate with it.
func boring(ctx context.Context, msg string) <-chan Message { // <-chan Message means receives-only channel of Message.
	c := make(chan Message)
	waitForIt := make(chan string) // share between all messages
	go func() {                    // we launch goroutine inside a function.
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println(msg, " exit!")
				return
			default:
			}

			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}

			fmt.Printf("send %s %d\n", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond) // 造成坑点

			// every time the goroutine send message.
			// This code waits until the value to be received.
			fmt.Println(<-waitForIt, "is ack!") // 造成坑点
		}
	}()
	return c // return a channel to caller.
}

/**
该模型有坑，waitForIt 可能被永久阻塞
*/
func main() {
	// 如果不调用cancel()会造成数据丢失
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel() // 进程安全退出
		fmt.Println("exec cancel()")
		time.Sleep(3 * time.Second)
	}()

	// merge 2 channels into 1 channel
	c := fanIn(ctx, boring(ctx, "Joe"))

	for i := 0; i < 6; i++ {
		msg1 := <-c // wait to receive message
		fmt.Println(msg1.str)

		// each go routine have to wait
		msg1.wait <- msg1.str // main goroutine allows the boring goroutine to send next value to message channel.
	}

	fmt.Println("You're both boring. I'm leaving")
}

// main: goroutine                                          boring: goroutine
//    |                                                           |
//    |                                                           |
// wait for receiving msg from channel c                    c <- Message{} // send message
//   <-c                                                          |
//    |                                                           |
//    |                                                     <-waitForIt // wait for wake up signal
// send value to channel                                          |
// hey, boring. You can send next value to me                     |
//   wait <-true                                                  |
///                            REPEAT THE PROCESS
