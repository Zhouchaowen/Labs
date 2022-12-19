// Test the performance difference of slice passing by value vs passing by pointer

// refer to:
// https://cloud.tencent.com/developer/article/1469185
// https://chende.ren/2021/01/06213457-012-benchmark-test.html

// go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
// go test -run none -bench . -benchtime 3s -benchmem

// pkg: Labs/lab_language/benchmarks/ch_6
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkPointer-4      1000000000               0.4151 ns/op          0 B/op          0 allocs/op
// BenchmarkValue-4        1000000000               0.4105 ns/op          0 B/op          0 allocs/op
package ch_6

import (
	"testing"
)

type Log struct {
	a string
	b string
	c string
	d string
}

func StructPointer(data *[]Log) {
	_ = data
}

func StructValue(data []Log) {
	_ = data
}

// BenchmarkPointer runs the test when the branch is predictable.
func BenchmarkPointer(b *testing.B) {
	data := make([]Log, 100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructPointer(&data)
	}
}

// BenchmarkValue runs the test when the branch is random.
func BenchmarkValue(b *testing.B) {
	data := make([]Log, 100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructValue(data)
	}
}
