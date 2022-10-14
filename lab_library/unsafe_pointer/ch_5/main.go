package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "ax~_"

	// []byte(s)
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	p := (*byte)(unsafe.Pointer(h.Data))
	b := unsafe.Slice(p, h.Len)

	// string(b)
	s2 := *(*string)(unsafe.Pointer(&b)) // string等于[]ArbitraryType转换

	fmt.Printf("string value: %s\n", s)
	fmt.Printf("string pointer: %p\n", &s)
	fmt.Printf("string pointer: %p\n", &s2)
	fmt.Println("----------------")
	fmt.Printf("StringHeader pointer: %p\n", h)
	fmt.Printf("StringHeader data pointer: %p\n", unsafe.Pointer(h.Data))
	fmt.Printf("StringHeader len  value: %d\n", h.Len)
	fmt.Println("----------------")
	fmt.Printf("StringHeader data pointer: %p\n", p)
	fmt.Printf("StringHeader data value[0]: %c\n", *p)
	fmt.Printf("StringHeader data value[1]: %c\n", *(*byte)(unsafe.Pointer(h.Data + uintptr(1))))
	fmt.Printf("StringHeader data value[2]: %c\n", *(*byte)(unsafe.Pointer(h.Data + uintptr(2))))
	fmt.Printf("StringHeader data value[3]: %c\n", *(*byte)(unsafe.Pointer(h.Data + uintptr(3))))
	fmt.Println("----------------")
	fmt.Printf("s2 value: %d\n", b)
}
