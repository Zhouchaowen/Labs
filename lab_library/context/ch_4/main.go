// 演示如何将一个值传递给上下文，以及如何在它存在时检索它。
package main

import (
	"context"
	"fmt"
)

// This example demonstrates how a value can be passed to the context
// and also how to retrieve it if it exists.
func ExampleWithValue() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

	// Output:
	// found value: Go
	// key not found: color
}

func main() {
	ExampleWithValue()
}

// WithValue如何保存值的
func ValueContext() {
	type orderID int
	var x = context.TODO()
	x = context.WithValue(x, orderID(1), "1234")
	x = context.WithValue(x, orderID(2), "2345")

	y := context.WithValue(x, orderID(3), "4567")
	x = context.WithValue(x, orderID(3), "3456")

	fmt.Println(x.Value(orderID(3)))
	fmt.Println(y.Value(orderID(3)))
}

/**
ctx 这么设计是为了能让代码每执行到一个点都可以根据当前情况嵌入新的上下文信息，但我们也可以看到，如果我们每次加一个新值都执行 WithValue 会导致 ctx 的树的层数过高，查找成本比较高 O(H)。
很多业务场景中，我们希望在请求入口存入值，在请求过程中随时取用。这时候我们可以将 value 作为一个 map 整体存入。
											  ┌────────────┐
											  │  emptyCtx  │
											  └────────────┘
													 ▲
													 │
													 │
													 │    parent
													 │
													 │
								   ┌───────────────────────────────────┐
								   │      valueCtx{k: 1, v: 1234}      │
								   └───────────────────────────────────┘
													 ▲
													 │
													 │
													 │    parent
													 │
													 │
													 │
								   ┌───────────────────────────────────┐
								   │      valueCtx{k: 2, v: 2345}      │
								   └───────────────────────────────────┘
													 ▲
													 │
								  ┌──────────────────┴──────────────────────┐
								  │                                         │
								  │                                         │
				┌───────────────────────────────────┐     ┌───────────────────────────────────┐
				│      valueCtx{k: 3, v: 3456}      │     │      valueCtx{k: 3, v: 4567}      │
				└───────────────────────────────────┘     └───────────────────────────────────┘
							 ┌───────┐                                   ┌───────┐
							 │   x   │                                   │   y   │
							 └───────┘                                   └───────┘
*/
