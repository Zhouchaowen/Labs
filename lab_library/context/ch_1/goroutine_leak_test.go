package main

import (
	"context"
	"go.uber.org/goleak"
	"testing"
)

// go test -v -run Leak$
// 测试 goroutine 泄露场景

func NoLeak() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done(): // 不返回可能发生泄露
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		if n == 5 {
			break
		}
	}
}

func Leak() {
	gen := func() <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	for n := range gen() {
		if n == 5 {
			break
		}
	}
}

func TestNoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	NoLeak()
}

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	Leak()
}

/**
产生goroutine leak的原因:
	goroutine由于channel的读/写端退出而一直阻塞，导致goroutine一直占用资源，而无法退出，如只有写入，没有接收，反之一样；
	goroutine进入死循环中，导致资源一直无法释放；
*/

// https://codeantenna.com/a/Q47HvbRdRm
// https://jasonkayzk.github.io/2021/04/21/%E4%BD%BF%E7%94%A8Uber%E5%BC%80%E6%BA%90%E7%9A%84goleak%E5%BA%93%E8%BF%9B%E8%A1%8Cgoroutine%E6%B3%84%E9%9C%B2%E6%A3%80%E6%B5%8B/
