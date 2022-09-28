// go test -run none -bench . -benchtime 3s -benchmem
// pkg: Labs/lab_benchmarks/ch_4
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkPredictable-4           4444155               785.5 ns/op             0 B/op          0 allocs/op
// BenchmarkUnpredictable-4         2322124              1439 ns/op               0 B/op          0 allocs/op

// 包预测提供代码来显示分支预测如何影响性能。
// Package prediction provides code to show how branch
// prediction can affect performance.
package prediction

import (
	"math/rand"
	"testing"
)

// crunch is used to perform branch instructions.
func crunch(data []uint8) uint8 {
	var sum uint8
	for _, v := range data {
		if v < 128 {
			sum--
		} else {
			sum++
		}
	}
	return sum
}

var fa uint8

// BenchmarkPredictable runs the test when the branch is predictable.
func BenchmarkPredictable(b *testing.B) {
	data := make([]uint8, 1024)
	b.ResetTimer()

	var a uint8

	for i := 0; i < b.N; i++ {
		a = crunch(data)
	}

	fa = a
}

// BenchmarkUnpredictable runs the test when the branch is random.
func BenchmarkUnpredictable(b *testing.B) {
	data := make([]uint8, 1024)
	rand.Seed(0)

	// Fill data with (pseudo) random noise
	for i := range data {
		data[i] = uint8(rand.Uint32())
	}

	b.ResetTimer()

	var a uint8

	for i := 0; i < b.N; i++ {
		a = crunch(data)
	}

	fa = a
}
