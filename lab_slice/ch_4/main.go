package main

import (
	"fmt"
	"unsafe"
)

func memoryAlignment1() {
	type t struct { // 对齐内存 12字节
		a int8
		b int32
		c int8
	}
	s := []t{{1, 2, 3}, {11, 22, 33}, {11, 22, 33}}
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryAlignment2() {
	type t struct { // 对齐内存 8字节
		a int8
		c int8
		b int32
	}
	s := []t{{1, 2, 3}, {11, 22, 33}, {11, 22, 33}}
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryAlignment3() {
	type t struct { // 对齐内存 32字节
		a int8
		c int8
		b []int32
	}
	s := make([]t, 3)
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct %+v , pointer %p\n", t{}, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func memoryAlignment4() {
	type t struct { // 对齐内存 40字节
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

func memoryAlignment5() {
	type t struct { // 对齐内存8字节
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

func memoryAlignment6() {
	type t struct { // 对齐内存8字节
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

func main() {
	memoryAlignment1()
	memoryAlignment2()
	memoryAlignment3()
	memoryAlignment4()
	memoryAlignment5()
	memoryAlignment6()
}
