package main

import (
	"fmt"
	"unsafe"
)

func memoryUsage1() {
	// |-----------------|
	// |-----|-----|-----|
	// |  a  |  b  |  c  |
	// |-----|-----|-----|
	// |  4  |  4  |  4  |
	// |-----|-----|-----|
	// |-----------------|
	type t struct { // 内存占用 12字节
		a int8  // 4
		b int32 // 4
		c int8  // 4
	}
	s := []t{{1, 2, 3}, {11, 22, 33}, {11, 22, 33}}
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryUsage2() {
	// |-----------|
	// |-----|-----|
	// | ab  |  c  |
	// |-----|-----|
	// |  4  |  4  |
	// |-----|-----|
	// |-----------|
	type t struct { // 内存占用 8字节
		a int8
		b int8
		c int32 // 4
	}
	s := []t{{1, 2, 3}, {11, 22, 33}, {11, 22, 33}}
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryUsage3() {
	// |-----------------------|
	// |-----|-----|-----|-----|
	// | ab  |  	  c		   |
	// |-----|-----|-----|-----|
	// |  8  |  8  |  8  |  8  |
	// |-----|-----|-----|-----|
	// |-----------------------|
	type t struct { // 内存占用 32字节
		a int8
		b int8
		c []int32
	}
	s := make([]t, 3)
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryUsage4() {
	type t struct { // 内存占用 40字节
		a int8
		c int64
		b []int32
	}
	s := make([]t, 3)
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryUsage5() {
	type t struct { // 内存占用 8字节
		a int8
		c int8
		b [1]int32
	}
	s := make([]t, 3)
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryUsage6() {
	type t struct { // 内存占用 8字节
		c int32
		a int16
		b [7]int8
	}
	s := make([]t, 3)
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryUsage7() {
	type t struct { // 内存占用 8字节
		a   uint64
		_p1 [8]uint64
		b   uint64
		_p2 [8]uint64
		c   uint64
		_p3 [8]uint64
	}
	s := make([]t, 3)
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func main() {
	memoryUsage1()
	memoryUsage2()
	memoryUsage3()
	memoryUsage4()
	memoryUsage5()
	memoryUsage6()
	memoryUsage7()
}
