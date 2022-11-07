// Tests to show the different performance based on concurrency with or without parallelism.

// GOGC=off GOMAXPROCS=1 go test -bench . -benchtime 3s

// Processing 500000000 numbers using 4 goroutines on 1 thread(s)
// pkg: Labs/lab_benchmarks/ch_3
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkSequential                    1        8184318940 ns/op
// BenchmarkConcurrent                    1        5186416376 ns/op
// BenchmarkSequentialAgain               8         412640644 ns/op
// BenchmarkConcurrentAgain               8         408666952 ns/op

// GOGC=off go test -bench . -benchtime 3s

// Processing 500000000 numbers using 4 goroutines on 4 thread(s)
// pkg: Labs/lab_benchmarks/ch_3
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkSequential-4                  1        5014648889 ns/op
// BenchmarkConcurrent-4                 14         242329434 ns/op
// BenchmarkSequentialAgain-4             7         458199166 ns/op
// BenchmarkConcurrentAgain-4            14         236224286 ns/op
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var numbers []int

func init() {
	rand.Seed(time.Now().UnixNano())
	numbers = generateList(5e8)
	fmt.Printf("Processing %d numbers using %d goroutines on %d thread(s)\n", len(numbers), runtime.NumCPU(), runtime.GOMAXPROCS(0))
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}

func BenchmarkSequentialAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrentAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}
