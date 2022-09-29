// 父ctx进行cancel()后，所有子ctx.Done()都收到信号。
package main

import (
	"context"
	"fmt"
	"time"
)

func funcA() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go funcB1(ctx)

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("cancel()")
	}
}

func funcB1(ctx context.Context) {
	go funcC1(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("funcB1 ", ctx.Err())
	}
}

func funcC1(ctx context.Context) {
	go funcD1(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("funcC1 ", ctx.Err())
	}
}

func funcD1(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("funcD1 ", ctx.Err())
	}
}

/*
实验 父ctx进行cancel()后，子ctx会怎样？
	1.父ctx进行cancel()后，所有子ctx.Done()都收到信号。
*/
func main() {
	funcA()

	time.Sleep(15 * time.Second)
}
