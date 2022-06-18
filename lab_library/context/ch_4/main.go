package main

import (
	"context"
	"fmt"
)

// This example demonstrates how a value can be passed to the context
// and also how to retrieve it if it exists.
// 演示如何将一个值传递给上下文，以及如何在它存在时检索它。
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
