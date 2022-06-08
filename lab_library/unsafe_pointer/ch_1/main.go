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
	s := []string{"abcd", "abc", "adca"}
	for i, _ := range s {
		fmt.Printf("%p\n", &s[i])
		fmt.Printf("%p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&s[i]))+uintptr(8))) // Len 地址
		fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[i])) + uintptr(8))))
	}
}
