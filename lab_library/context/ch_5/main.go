// 父context传递信息，子context继续传递消息
package main

import (
	"context"
	"fmt"
	"time"
)

func funcA() {
	ctx := context.WithValue(context.Background(), "KeyA", "ValueA")

	go funcB1(ctx)
	go funcB2(ctx)

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("time out")
	}
}

func funcB1(ctx context.Context) {
	v := ctx.Value("KeyA")
	fmt.Println("funcB1: ", v)
	ctx = context.WithValue(ctx, "KeyB", "ValueB1")

	go funcC1(ctx)
}

func funcB2(ctx context.Context) {
	v := ctx.Value("KeyA")
	fmt.Println("funcB2: ", v)

	ctx = context.WithValue(ctx, "KeyB", "ValueB2")
	go funcC2(ctx)

}

func funcC1(ctx context.Context) {
	vA := ctx.Value("KeyA")
	vB := ctx.Value("KeyB")
	fmt.Println("funcC1: ", vA)
	fmt.Println("funcC1: ", vB)
}

func funcC2(ctx context.Context) {
	vA := ctx.Value("KeyA")
	vB := ctx.Value("KeyB")
	fmt.Println("funcC2: ", vA)
	fmt.Println("funcC2: ", vB)
}

func main() {
	funcA()
}
