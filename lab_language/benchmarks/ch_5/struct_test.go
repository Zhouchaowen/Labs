// https://cloud.tencent.com/developer/article/1469185
// https://chende.ren/2021/01/06213457-012-benchmark-test.html
package ch_5

import (
	"math/rand"
	"runtime"
	"testing"
)

// go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
// go test -run none -bench . -benchtime 3s -benchmem
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
