package main

import (
	"testing"
)

func CallSlice(s []int) {

}
func CallArr(s [10000]int) {

}

func BenchmarkCallSlice(b *testing.B) {
	s := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		CallSlice(s)
	}
}

func BenchmarkCallArr(b *testing.B) {
	var a [10000]int
	for i := 0; i < b.N; i++ {
		CallArr(a)
	}
}

// go test -gcflags "-N -l" -bench='^Benchmark' -benchmem .
