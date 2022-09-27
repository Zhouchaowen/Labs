package main

import (
	"fmt"
	"unsafe"
)

/*
实验：string由两部分组成
type StringHeader struct {
	Data uintptr
	Len  int
}
*/
func main() {
	s := []string{"abcd", "abc", "adcaadcaadcaadcaadcaadcaadcaadca"}
	fmt.Printf("s arr pointer: %p\n", &s)
	fmt.Printf("s arr data pointer: %x\n", *(*uintptr)(unsafe.Pointer(&s))) // 涉及slice结构

	fmt.Printf("s arr data pointer: %p\n", &s[0])
	fmt.Printf("string pointer: %p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&s[0]))+uintptr(8))) // Len 地址
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(8))))             // Len 值
	fmt.Println(*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])))))                       // string 值

	fmt.Printf("%p\n", &s[1])
	fmt.Printf("%p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&s[0]))+uintptr(16)))        // Len 地址
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(16+8))))  // Len 值
	fmt.Println(*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(16)))) // string 值

	fmt.Printf("%p\n", &s[2])
	fmt.Printf("%p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&s[0]))+uintptr(32)))        // Len 地址
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(32+8))))  // Len 值
	fmt.Println(*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(32)))) // string 值
}
