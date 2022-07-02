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
	case <-ctx.Done():
		fmt.Println("funcA ", ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

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
实验 子ctx被cancel()后父ctx会怎样？
	1.子ctx被cancel()，子ctx的所有孙子ctx.Done()都收到信号。
	2.子ctx的所有父亲不受影响，ctx.Done()不会收到信号。
*/
func main() {
	funcA()

	time.Sleep(15 * time.Second)
}
