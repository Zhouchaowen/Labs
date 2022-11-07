// the performance gap of concat/join/buffer concatenating strings

// Building strings dynamically can cause performance problems.
// The addition operator concatenates strings, reallocating memory and copying data each time.
// An improvement is to preallocate memory and return it all at once.

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

// BenchmarkConcat tests the performance of using concat.
func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !concat() {
			b.Fatal()
		}
	}
}

// BenchmarkJoin tests the performance of using join.
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !join() {
			b.Fatal()
		}
	}
}

// BenchmarkBuffer tests the performance of using buffer.
func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !buffer() {
			b.Fatal()
		}
	}
}
