package main

import (
	"fmt"
	"unsafe"
)

/*
实验：slice由三部分组成
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
*/
func main() {
	arr := make([]string, 0, 4)
	arr = append(arr, "abcd")

	fmt.Printf("arr addr: %p\n", &arr)

	fmt.Printf("arr len addr: %p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&arr))+uintptr(8))) // Len 地址
	fmt.Printf("arr len: %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + uintptr(8))))

	fmt.Printf("arr cap addr: %p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&arr))+uintptr(16))) // Cap 地址
	fmt.Printf("arr cap: %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + uintptr(16))))
}
