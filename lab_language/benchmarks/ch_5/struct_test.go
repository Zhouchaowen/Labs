// Test the performance difference of struct passing by value vs passing by pointer

// refer to:
// https://cloud.tencent.com/developer/article/1469185
// https://chende.ren/2021/01/06213457-012-benchmark-test.html

// go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
// go test -run none -bench . -benchtime 3s -benchmem

// pkg: Labs/lab_language/benchmarks/ch_5
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkPointer-4      120547652               30.45 ns/op            0 B/op          0 allocs/op
// BenchmarkValue-4        115684429               31.13 ns/op            0 B/op          0 allocs/op
package ch_5

import (
	"math/rand"
	"runtime"
	"testing"
)

type Log struct {
	a string
	b string
	c string
	d string
}

func StructPointer(data *Log) {
	_ = data
}

func StructValue(data Log) {
	_ = data
}

func genDataPointer(n int) []*Log {
	data := make([]*Log, n)
	for i := 0; i < n; i++ {
		data[i] = &Log{
			a: "a",
			b: "b",
			c: "c",
			d: "d",
		}
	}
	return data
}

func genDataValue(n int) []Log {
	data := make([]Log, n)
	for i := 0; i < n; i++ {
		data[i] = Log{
			a: "a",
			b: "b",
			c: "c",
			d: "d",
		}
	}
	return data
}

// BenchmarkPointer
func BenchmarkPointer(b *testing.B) {
	var data = genDataPointer(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructPointer(data[rand.Intn(100)])
	}
	runtime.GC()
}

// BenchmarkValue
func BenchmarkValue(b *testing.B) {
	var data = genDataValue(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructValue(data[rand.Intn(100)])
	}
	runtime.GC()
}
