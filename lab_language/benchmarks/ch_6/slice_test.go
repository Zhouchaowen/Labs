// https://cloud.tencent.com/developer/article/1469185
// https://chende.ren/2021/01/06213457-012-benchmark-test.html
package ch_5

import (
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

func StructPointer(data *[]Log) {
	_ = data
}

func StructValue(data []Log) {
	_ = data
}

// BenchmarkPredictable runs the test when the branch is predictable.
func BenchmarkPointer(b *testing.B) {
	data := make([]Log, 100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructPointer(&data)
	}
}

// BenchmarkUnpredictable runs the test when the branch is random.
func BenchmarkValue(b *testing.B) {
	data := make([]Log, 100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructValue(data)
	}
}
