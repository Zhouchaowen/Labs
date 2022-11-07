// Test the performance difference between strong conversion and standard conversion

// refer to:
// https://segmentfault.com/a/1190000037679588

// go test -run none -bench . -benchtime 3s -benchmem

// pkg: Labs/lab_language/benchmarks/ch_7
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// Benchmark_NormalBytes2String-4          78036273                41.32 ns/op           48 B/op          1 allocs/op
// Benchmark_Byte2String-4                 1000000000               0.4212 ns/op          0 B/op          0 allocs/op
// Benchmark_NormalString2Bytes-4          66684045                52.62 ns/op           48 B/op          1 allocs/op
// Benchmark_String2Bytes-4                1000000000               0.4196 ns/op          0 B/op          0 allocs/op
package ch_7

import (
	"bytes"
	"reflect"
	"testing"
	"unsafe"
)

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// TestBytes2String Test the strong conversion function
func TestBytes2String(t *testing.T) {
	x := []byte("Hello Gopher!")
	y := Bytes2String(x)
	z := string(x)

	if y != z {
		t.Fail()
	}
}

// TestString2Bytes Test the strong conversion function
func TestString2Bytes(t *testing.T) {
	x := "Hello Gopher!"
	y := String2Bytes(x)
	z := []byte(x)

	if !bytes.Equal(y, z) {
		t.Fail()
	}
}

// Benchmark_NormalBytes2String Test standard conversion string() performance
func Benchmark_NormalBytes2String(b *testing.B) {
	x := []byte("Hello Gopher! Hello Gopher! Hello Gopher!")
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

// 测试强转换[]byte到string性能
func Benchmark_Byte2String(b *testing.B) {
	x := []byte("Hello Gopher! Hello Gopher! Hello Gopher!")
	for i := 0; i < b.N; i++ {
		_ = Bytes2String(x)
	}
}

// Benchmark_NormalString2Bytes Test standard conversion []byte performance
func Benchmark_NormalString2Bytes(b *testing.B) {
	x := "Hello Gopher! Hello Gopher! Hello Gopher!"
	for i := 0; i < b.N; i++ {
		_ = []byte(x)
	}
}

// Benchmark_String2Bytes Test cast string to []byte performance
func Benchmark_String2Bytes(b *testing.B) {
	x := "Hello Gopher! Hello Gopher! Hello Gopher!"
	for i := 0; i < b.N; i++ {
		_ = String2Bytes(x)
	}
}
