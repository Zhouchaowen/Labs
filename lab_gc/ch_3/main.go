package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

const N = 30e6

func mapInt32PointerTest() {
	// Big map with a pointer in the value
	m := make(map[int32]*int32)
	for i := 0; i < N; i++ {
		n := int32(i)
		m[n] = &n
	}
	runtime.GC()
	fmt.Printf("With %T, GC took %s\n", m, timeGC())
	_ = m[0] // Preserve m until here, hopefully
}

func mapNInt32PointerTest() {
	// Big map (preallocated) with a pointer in the value
	m := make(map[int32]*int32, N)
	for i := 0; i < N; i++ {
		n := int32(i)
		m[n] = &n
	}
	runtime.GC()
	fmt.Printf("With %T, GC took %s\n", m, timeGC())
	_ = m[0] // Preserve m until here, hopefully
}

func mapInt32Test() {
	// Big map, no pointer in the value
	m := make(map[int32]int32)
	for i := 0; i < N; i++ {
		n := int32(i)
		m[n] = n
	}
	runtime.GC()
	fmt.Printf("With %T, GC took %s\n", m, timeGC())
	_ = m[0]
}

func mapNInt32Test() {
	// Big map, no pointer in the value, map fully pre-allocated
	m := make(map[int32]int32, N)
	for i := 0; i < N; i++ {
		m[int32(i)] = int32(i)
	}
	runtime.GC()
	fmt.Printf("With %T, GC took %s\n", m, timeGC())
	_ = m[0]
}

func mapStringTest() {
	// Big map with a pointer in the key
	m := make(map[string]string)
	for i := 0; i < N; i++ {
		m[strconv.Itoa(i)] = strconv.Itoa(i)
	}
	runtime.GC()
	fmt.Printf("With %T, GC took %s\n", m, timeGC())
	_ = m["0"]
}

func mapNStringTest() {
	// Big map with a pointer in the key, map fully pre-allocated
	m := make(map[string]string, N)
	for i := 0; i < N; i++ {
		m[strconv.Itoa(i)] = strconv.Itoa(i)
	}
	runtime.GC()
	fmt.Printf("With %T, GC took %s\n", m, timeGC())
	_ = m["0"]
}

func structInt32Test() {
	// A slice, just for comparison to show that
	// merely holding onto millions of int32s is fine
	// if they're in a slice.
	type t struct {
		p, q int32
	}
	var s []t
	for i := 0; i < N; i++ {
		n := int32(i)
		s = append(s, t{n, n})
	}
	runtime.GC()
	fmt.Printf("With a plain slice (%T), GC took %s\n", s, timeGC())
	_ = s[0]
}

func structNInt32Test() {
	// A slice, just for comparison to show that
	// merely holding onto millions of int32s is fine
	// if they're in a pre-allocated slice.
	type t struct {
		p, q int32
	}
	s := make([]t, 0, N)
	for i := 0; i < N; i++ {
		s = append(s, t{int32(i), int32(i)})
	}
	runtime.GC()
	fmt.Printf("With a plain slice (%T), GC took %s\n", s, timeGC())
	_ = s[0]
}

func sliceStringTest() {
	var s []string
	for i := 0; i < N; i++ {
		s = append(s, strconv.Itoa(i))
	}
	runtime.GC()
	fmt.Printf("With a plain slice (%T), GC took %s\n", s, timeGC())
	_ = s[0]
}

func sliceNStringTest() {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = strconv.Itoa(i)
	}
	runtime.GC()
	fmt.Printf("With a plain slice (%T), GC took %s\n", s, timeGC())
	_ = s[0]
}

func main() {
	mapInt32PointerTest()
	mapNInt32PointerTest()
	mapInt32Test()
	mapNInt32Test()
	mapStringTest()
	mapNStringTest()
	structInt32Test()
	structNInt32Test()
	sliceStringTest()
	sliceNStringTest()
}

func timeGC() time.Duration {
	start := time.Now()
	runtime.GC()
	return time.Since(start)
}
