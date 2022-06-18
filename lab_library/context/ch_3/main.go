package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond // a reasonable duration to block in an example

// This example passes a context with a timeout to tell a blocking function that
// it should abandon its work after the timeout elapses.
// 演示超时的上下文来告诉一个阻塞函数它应该在超时后放弃它的工作.
func ExampleWithTimeout() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

	// Output:
	// context deadline exceeded
}

func main() {
	ExampleWithTimeout()
}
