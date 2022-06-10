package main

import (
	"sync/atomic"
	"testing"
)

// go test -gcflags "-N -l" -bench='Increase$' -benchmem .

/*

在并发编程中，经常会有共享数据被多个goroutine同时访问， 所以如何有效的进行数据的设计，就是一个相当有技巧的操作。
最常用的技巧就是Padding。现在大部分的CPU的cahceline是64字节，将变量补足为64字节可以保证它正好可以填充一个cacheline。

*/

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (np *NoPad) Increase() {
	atomic.AddUint64(&np.a, 1)
	atomic.AddUint64(&np.b, 1)
	atomic.AddUint64(&np.c, 1)
}

type Pad struct {
	a   uint64
	_p1 [8]uint64
	b   uint64
	_p2 [8]uint64
	c   uint64
	_p3 [8]uint64
}

func (p *Pad) Increase() {
	atomic.AddUint64(&p.a, 1)
	atomic.AddUint64(&p.b, 1)
	atomic.AddUint64(&p.c, 1)
}

func BenchmarkPad_Increase(b *testing.B) {
	pad := &Pad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pad.Increase()
		}
	})
}

func BenchmarkNoPad_Increase(b *testing.B) {
	nopad := &NoPad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			nopad.Increase()
		}
	})
}
