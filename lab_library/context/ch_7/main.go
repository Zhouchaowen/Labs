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
	ctx = context.WithValue(ctx, "KeyC", "ValueC")
	go funcD1(ctx)
}

func funcD1(ctx context.Context) {
	vC := ctx.Value("KeyC")
	fmt.Println("funcC1: ", vC)
	select {
	case <-ctx.Done():
		fmt.Println("funcD1 ", ctx.Err())
	}
}

/*
实验 子ctx被改变类型后，父ctx进行cancel()会怎样？
	1.子ctx被改变类型后，父ctx进行cancel()，所有子ctx.Done()都收到信号。
*/
func main() {
	funcA()

	time.Sleep(15 * time.Second)
}
