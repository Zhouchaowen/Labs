// go test -run none -bench . -benchtime 3s -benchmem

// pkg: Labs/lab_benchmarks/ch_1
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkSprint-4       34773686               104.1 ns/op             5 B/op          1 allocs/op
// BenchmarkSprintf-4      42649080                84.62 ns/op            5 B/op          1 allocs/op

// Sprint/Sprintf的性能差距
// Basic benchmark test.
package ch_1

import (
	"fmt"
	"testing"
)

var gs string

// BenchmarkSprint tests the performance of using Sprint.
func BenchmarkSprint(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	gs = s
}

// BenchmarkSprint tests the performance of using Sprintf.
func BenchmarkSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}

	gs = s
}
