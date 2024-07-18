// go test -bench . -benchmem -memprofile p.out -gcflags -m=2
// go tool pprof -noinlines p.out
// https://github.com/ardanlabs/gotraining/blob/6b67abeed7a8875a2faf7eb0cf4e67e6b1fb2ff3/topics/go/profiling/memcpu/README.md
// Tests to see how each algorithm compare.
package main

import (
	"bytes"
	"testing"
)

var output bytes.Buffer
var in = assembleInputStream()
var find = []byte("elvis")
var repl = []byte("Elvis")

// Capture the time it takes to execute algorithm one.
func BenchmarkAlgorithmOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkAlgorithmTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algTwo(in, find, repl, &output)
	}
}
