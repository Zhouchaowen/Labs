// Test the performance difference between strong conversion and standard conversion

// go test -run none -bench . -benchtime 3s -benchmem

// pkg: Labs/lab_language/benchmarks/ch_8
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// Benchmark_ChannelBigData-4       2664240              1291 ns/op               0 B/op          0 allocs/op
// Benchmark_ChannelSmallData-4    20802163               158.6 ns/op             8 B/op          1 allocs/op
package ch_8

import (
	"testing"
)

// Benchmark_ChannelBigData
func Benchmark_ChannelBigData(b *testing.B) {
	ch := make(chan [1000]int, 10)
	go func() {
		for {
			<-ch
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- [1000]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
}

// Benchmark_ChannelSmallData
func Benchmark_ChannelSmallData(b *testing.B) {
	ch := make(chan []int, 10)
	go func() {
		for {
			<-ch
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- []int{1}
	}
}
