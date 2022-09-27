package main

import (
	"fmt"
	"unsafe"
)

/*
实验：slice由三部分组成
type SliceHeader struct {
	Data uintptr	// 存储的就是一个指针
	Len  int
	Cap  int
}
*/
func main() {
	arr := make([]string, 0, 4)
	arr = append(arr, "abcd")
	arr = append(arr, "abcdc")
	arr = append(arr, "abcdce")

	fmt.Printf("arr addr: %p\n", &arr)

	fmt.Printf("arr len addr: %p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&arr))+uintptr(8))) // Len 地址
	fmt.Printf("arr len: %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + uintptr(8))))

	fmt.Printf("arr cap addr: %p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&arr))+uintptr(16))) // Cap 地址
	fmt.Printf("arr cap: %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + uintptr(16))))

	// 把arr地址取出来，然后强转为一个指针地址类型(*uintptr)，然后在取这个指针地址类型存储的值(值也是一个指针)
	fmt.Printf("arr addr uintptr: %x\n", *(*uintptr)(unsafe.Pointer(&arr))) // arr的[Data uintptr]存储的是一个地址
	fmt.Printf("arr addr uintptr: %p\n", &arr[0])                           // string 的指针

	fmt.Printf("arr addr[0]: %s\n", arr[0])                                                                        // string 值
	fmt.Printf("arr addr[1]: %v\n", *(*string)(unsafe.Pointer((*(*uintptr)(unsafe.Pointer(&arr))) + uintptr(16)))) // string 值
	fmt.Printf("arr addr[2]: %v\n", *(*string)(unsafe.Pointer((*(*uintptr)(unsafe.Pointer(&arr))) + uintptr(32)))) // string 值
}
