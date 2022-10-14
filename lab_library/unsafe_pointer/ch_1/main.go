// https://www.yuque.com/qyuhen/go/crac24
package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

/*
实验：string由两部分组成
type StringHeader struct {
	Data uintptr
	Len  int
}

借助 unsafe 实现指针转换和运算，须自行确保内存安全。
	1.普通指针： *int，包含类型信息。
	2.通用指针：Pointer，只有地址，没有类型。
	3.指针整数：uintptr，足以存储地址的整数。
普通指针（*int）和通用指针（Pointer）都能构成引用，影响垃圾回收。
而 uintptr 只是整数，不构成引用关系，无法阻止垃圾回收器清理目标对象。
*/
func testStringHeader1() {
	s := []string{"abcd", "abc", "adcaadcaadcaadcaadcaadcaadcaadca"}
	fmt.Printf("string arr addr pointer     : %p\n", &s)
	fmt.Printf("string arr data addr pointer: 0x%x\n", *(*uintptr)(unsafe.Pointer(&s))) // 涉及slice结构
	fmt.Println("-----------------------")
	fmt.Printf("string arr[0] data pointer: %p\n", &s[0])
	fmt.Printf("string arr[0] data value  : %s\n", *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])))))           // string 值
	fmt.Printf("string arr[0] len  pointer: %p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&s[0]))+uintptr(8)))            // Len 地址
	fmt.Printf("string arr[0] len  value  : %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(8)))) // Len 值
	fmt.Println("-----------------------")
	fmt.Printf("string arr[1] data pointer: %p\n", &s[1])
	fmt.Printf("string arr[1] data value  : %s\n", *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(16)))) // string 值
	fmt.Printf("string arr[1] len  pointer: %p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&s[0]))+uintptr(16+8)))             // Len 地址
	fmt.Printf("string arr[1] len  value  : %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(16+8))))  // Len 值
	fmt.Println("-----------------------")
	fmt.Printf("s arr[2] data pointer: %p\n", &s[2])
	fmt.Printf("s arr[2] data value  : %s\n", *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(32)))) // string 值
	fmt.Printf("s arr[2] len  pointer: %p\n", unsafe.Pointer(uintptr(unsafe.Pointer(&s[0]))+uintptr(32+8)))             // Len 地址
	fmt.Printf("s arr[2] len  value  : %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + uintptr(32+8))))  // Len 值
}

func testStringHeader2() {
	str := "hello,world"
	arr := str[:4]
	p1 := (*reflect.StringHeader)(unsafe.Pointer(&str))
	p2 := (*reflect.StringHeader)(unsafe.Pointer(&arr))

	fmt.Printf("%#v\n%#v\n", p1, p2)
}

// 要修改字符串，须转换为可变类型（[]rune 或 []byte），待完成后再转换回来。
// 但不管如何转换，都需重新分配内存，并复制数据。
func testStringHeader3() {
	s := strings.Repeat("a", 1<<10)

	// 分配内存、复制。
	bs := []byte(s)
	bs[1] = 'B'

	// 分配内存、复制。
	s2 := string(bs)

	hs := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hbs := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	hs2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("%#v\n%#v\n%#v\n", hs, hbs, hs2)
}

/*
 +-----------+          +---+---+---+---+---+
 |  pointer -|--------> | h | e | l | l | o |
 +-----------+          +---+---+---+---+---+
 |  len = 5  |
 +-----------+          [...]byte, UTF-8

    header
*/
func main() {
	fmt.Println("-----------------------testStringHeader1-----------------------")
	testStringHeader1()
	fmt.Println("-----------------------testStringHeader2-----------------------")
	testStringHeader2()
	fmt.Println("-----------------------testStringHeader3-----------------------")
	testStringHeader3()
}
