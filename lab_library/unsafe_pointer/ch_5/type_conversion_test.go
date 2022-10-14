package main

import (
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

var S = strings.Repeat("a", 100)

func normalConv() bool {
	b := []byte(S)
	s2 := string(b)
	return s2 == S
}

func unsafeConv() bool {
	// []byte(s)
	h := (*reflect.StringHeader)(unsafe.Pointer(&S))
	p := (*byte)(unsafe.Pointer(h.Data))
	b := unsafe.Slice(p, h.Len)

	// string(b)
	s2 := *(*string)(unsafe.Pointer(&b))

	return s2 == S
}

// ----------------------------------
// 普通转换 normalConv 调用  runtime.stringtoslicebyte、runtime.slicebytetostring，
// 引发 mallocgc、memmove 等操作。

// go test -bench .benchmem -v
func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !normalConv() {
			b.Fatal()
		}
	}
}

func BenchmarkUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !unsafeConv() {
			b.Fatal()
		}
	}
}
