package main

import (
	"fmt"
	"time"
	"unsafe"
)

/**
[] 24个字节
[3]byte 3个字节
[3]int32 12个字节
*  8 个字节
string 16个人字节
interface 16个人字节
func 8 个字节
*/

func StructInt() {
	type t struct { // 8 字节
		a int8
		b int32
	}
	s := []t{{1, 2}, {11, 22}, {111, 222}}
	fmt.Printf("StructInt slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructInt slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructString() {
	type t struct { // 16 字节
		a string
	}
	s := []t{{"1"}, {"2"}, {"3"}}
	fmt.Printf("StructString slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructString slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructStringPointer() {
	type t struct { // 8 字节
		a *string
	}
	s := []t{{new(string)}, {new(string)}, {new(string)}}
	fmt.Printf("StructPointer slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructPointer slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructInterface() {
	type t struct { // 16 字节
		a interface{}
	}
	s := make([]t, 3)
	fmt.Printf("StructInterface slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructInterface slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

type inter interface {
	now() int8
}

func StructFuncInterface() {
	type tt struct { // 16 字节
		a inter
	}
	s := make([]tt, 3)
	fmt.Printf("StructFuncInterface slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructFuncInterface slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(tt{}))
}

type tmp struct {
	b int64
}

func (t tmp) now() int8 {
	return 0
}

func StructFuncImplementInterface() {
	type tt struct { // 16 字节
		a inter
	}
	s := []tt{{tmp{1}}, {tmp{2}}, {tmp{3}}}
	fmt.Printf("StructFuncImplementInterface slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructFuncImplementInterface slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(tt{}))
}

func StructFunc() {
	type op func(a, b int8) int

	type t struct { // 8 字节
		a op
	}

	s := make([]t, 3)
	fmt.Printf("StructFunc slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructFunc slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructSliceByte() {
	type t struct { //  24 字节
		a []byte
	}
	s := []t{{[]byte{'a', 'b', 'c'}}, {[]byte{'d', 'e', 'f'}}, {[]byte{'g', 'h', 'i'}}}
	fmt.Printf("StructSliceByte slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructSliceByte slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructSliceByteN() {
	type t struct { // 3 字节
		a [3]byte
	}
	s := []t{{[3]byte{'a', 'b', 'c'}}, {[3]byte{'d', 'e', 'f'}}, {[3]byte{'g', 'h', 'i'}}}
	fmt.Printf("StructSliceByteN slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructSliceByteN slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructSliceInt32() {
	type t struct { // 24 字节
		a []int32
	}
	s := []t{{[]int32{1, 2, 3}}, {[]int32{4, 5, 6}}, {[]int32{7, 8, 9}}}
	fmt.Printf("StructSliceInt32 slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructSliceInt32 slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructSliceInt32N() {
	type t struct { // 12 字节
		a [3]int32
	}
	s := []t{{[3]int32{1, 2, 3}}, {[3]int32{4, 5, 6}}, {[3]int32{7, 8, 9}}}
	fmt.Printf("StructSliceInt32N slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructSliceInt32N slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructSliceString() {
	type t struct { // 24 字节
		a []string
	}
	s := []t{{[]string{"1", "2", "3"}}, {[]string{"4", "5", "6"}}, {[]string{"7", "8", "9"}}}
	fmt.Printf("StructSliceString slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructSliceString slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructSliceStringN() {
	type t struct { // 48 字节
		a [3]string
	}
	s := []t{{[3]string{"1", "2", "3"}}, {[3]string{"4", "5", "6"}}, {[3]string{"7", "8", "9"}}}
	fmt.Printf("StructSliceStringN slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructSliceStringN slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructSliceInt8Pointer() {
	type t struct { // 24 字节
		a []*int8
	}
	s := make([]t, 3)
	fmt.Printf("StructSliceInt8Pointer slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructSliceInt8Pointer slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructSliceStringPointer() {
	type t struct { // 24 字节
		a []*string
	}
	s := make([]t, 3)
	fmt.Printf("StructSliceStringPointer slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructSliceStringPointer slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructPointerSliceString() {
	type t struct { // 8 字节
		a *[]string
	}
	s := make([]t, 3)
	fmt.Printf("StructPointerSliceString slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructPointerSliceString slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructStruct() {
	type t struct { // 3 字节
		a int8
		b int8
		c int8
	}

	type tt struct { // 3 字节
		a t
	}
	s := make([]tt, 3)
	fmt.Printf("StructStruct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructStruct slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(tt{}))
}

func StructStructPointer() {
	type t struct { // 3 字节
		a int8
		b int8
		c int8
	}

	type tt struct { // 8 字节
		a *t
	}
	s := make([]tt, 3)
	fmt.Printf("StructStructPointer slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructStructPointer slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(tt{}))
}

func StructMapString() {
	type t struct { // 8 字节
		a map[string]string
	}
	s := make([]t, 3)
	fmt.Printf("StructMapString slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructMapString slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructMapPointerString() {
	type t struct { // 8 字节
		a *map[string]string
	}
	s := make([]t, 3)
	fmt.Printf("StructMapPointerString slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructMapPointerString slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructTimeString() {
	type t struct { // 8 字节
		a time.Time
	}
	s := make([]t, 3)
	fmt.Printf("StructTimeString slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructTimeString slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func StructTimePointerString() {
	type t struct { // 8 字节
		a *time.Time
	}
	s := make([]t, 3)
	fmt.Printf("StructTimePointerString slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("StructTimePointerString slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func main() {
	StructInt()           // 8 字节
	StructString()        // 16 字节
	StructStringPointer() // 8 字节

	StructInterface()              // 16 字节
	StructFuncInterface()          // 16 字节
	StructFuncImplementInterface() // 16 字节
	StructFunc()                   // 8 字节

	StructSliceByte()    // 24 字节
	StructSliceByteN()   // n*byte 字节
	StructSliceInt32()   // 24 字节
	StructSliceInt32N()  // n*int 字节
	StructSliceString()  // 24 字节
	StructSliceStringN() // n*string 字节

	StructSliceInt8Pointer()   // 24 字节
	StructSliceStringPointer() // 24 字节
	StructPointerSliceString() // 8 字节
	StructStruct()             // 3 字节
	StructStructPointer()      // 8 字节

	StructMapString()        // 8 字节
	StructMapPointerString() // 8 字节

	StructTimeString()        // 24 字节
	StructTimePointerString() // 8 字节
}
