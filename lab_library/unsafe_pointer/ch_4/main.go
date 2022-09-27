package main

import (
	"fmt"
	"unsafe"
)

/*
实验：struct的组成
 内存对齐
*/

type A struct {
	a uint8
	b int32
	c int64
}

func main() {
	a := A{1, 11, 1111}

	fmt.Printf("A addr: %p\n", &a)
	fmt.Printf("A.a offset: %d\n", unsafe.Offsetof(a.a))
	fmt.Printf("A.b offset: %d\n", unsafe.Offsetof(a.b))
	fmt.Printf("A.c offset: %d\n", unsafe.Offsetof(a.c))

	// 强行改变变量类型
	aa := (*int32)(unsafe.Pointer(&a)) // 修改A.a的值
	*aa = 111111

	b := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + uintptr(4))) // 修改A.b的值
	*b = 22

	fmt.Printf("A.a: %d\n", *(*uint8)(unsafe.Pointer(&a)))
	fmt.Printf("A.b: %d\n", *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + unsafe.Offsetof(a.b))))
	fmt.Printf("A.c: %d\n", *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + unsafe.Offsetof(a.c))))
}
