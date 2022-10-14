// 动态构建字符串也容易造成性能问题。
// 加法操作符拼接字符串，每次都需重新分配内存和复制数据。
// 改进方法是预分配内存，然后一次性返回。
// go test -bench . -benchmem -v

// pkg: Labs/lab_language/benchmarks/ch_1
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkConcat-4           5604            229388 ns/op          530274 B/op        999 allocs/op
// BenchmarkJoin-4            83162             13194 ns/op            1024 B/op          1 allocs/op
// BenchmarkBuffer-4         144687              8485 ns/op            2048 B/op          2 allocs/op
package ch_1

import (
	"bytes"
	"strings"
	"testing"
)

const N = 1000
const C = "a"

var S = strings.Repeat(C, N)

func concat() bool {
	var s2 string
	for i := 0; i < N; i++ {
		s2 += C
	}
	return s2 == S
}

func join() bool {
	b := make([]string, N)
	for i := 0; i < N; i++ {
		b[i] = C
	}

	return strings.Join(b, "") == S
}

func buffer() bool {
	var b bytes.Buffer
	b.Grow(N)

	for i := 0; i < N; i++ {
		b.WriteString(C)
	}

	return b.String() == S
}

// ----------------------------------

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !concat() {
			b.Fatal()
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !join() {
			b.Fatal()
		}
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !buffer() {
			b.Fatal()
		}
	}
}
