package main

import (
	"fmt"
	"unsafe"
)

/*
type               size      default
------------------ --------- ----------- --------------------------
bool               1         false
byte               1         0			uint8

int,   uint      4 or 8      0
int8,  uint8       1         0          -128 ~ 127, 0 ~ 255
int16, uint16      2         0          -32768 ~ 32767, 0 ~ 65535
int32, uint32      4         0
int64, uint64      8         0

float32            4         0.0
float64            8         0.0

complex64          8
complex128        16

rune               4         0			unicode code point, int32
uintptr          4 or 8      0			uint

string                       ""
array
struct
function                     nil
interface                    nil
map                          nil
slice                        nil
channel                      nil
*/
func uint8Pointer() {
	s := []uint8{1, 2, 33, 44}
	fmt.Printf("uint8 slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("uint8 slice idx %d , pointer %p\n", i, &s[i])
	}
}

func int8Pointer() {
	s := []int8{1, 2, 33, 44}
	fmt.Printf("int8 slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("int8 slice idx %d , pointer %p\n", i, &s[i])
	}
}

func int16Pointer() {
	s := []int16{1, 2, 333, 444}
	fmt.Printf("uint16 slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("uint16 slice idx %d , pointer %p\n", i, &s[i])
	}
}

func int32Pointer() {
	s := []int32{1, 2, 333, 444}
	fmt.Printf("int32 slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("int32 slice idx %d , pointer %p\n", i, &s[i])
	}
}

func int64Pointer() {
	s := []int64{1, 2, 333, 444}
	fmt.Printf("int64 slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("int64 slice idx %d , pointer %p\n", i, &s[i])
	}
}

func intPointer() {
	s := []int64{1, 2, 333, 444}
	fmt.Printf("int slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("int slice idx %d , pointer %p\n", i, &s[i])
	}
}

func float32Pointer() {
	s := []float32{1.0, 2.4, 333.131, 444}
	fmt.Printf("float32 slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("float32 slice idx %d , pointer %p\n", i, &s[i])
	}
}

func float64Pointer() {
	s := []float64{1.0, 2.4, 333.131, 444}
	fmt.Printf("float64 slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("float64 slice idx %d , pointer %p\n", i, &s[i])
	}
}

func boolPointer() {
	s := []bool{true, false, true}
	fmt.Printf("bool slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("bool slice idx %d , pointer %p\n", i, &s[i])
	}
}

func stringPointer() {
	s := []string{"1", "22", "333", "444"}
	fmt.Printf("string slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("string slice idx %d , pointer %p\n", i, &s[i])
	}
}

func StructPointer() {
	type t struct {
		a int8
		b int32
	}
	s := []t{{1, 2}, {11, 22}, {111, 222}}
	fmt.Printf("Struct slice pointer %p\n", s)
	for i, _ := range s {
		fmt.Printf("Struct slice idx %d , pointer %p\n", i, &s[i])
	}

	fmt.Println(unsafe.Sizeof(t{}))
}

func main() {
	uint8Pointer()   // 1 字节
	int8Pointer()    // 1 字节
	int16Pointer()   // 2 字节
	int32Pointer()   // 4 字节
	int64Pointer()   // 8 字节
	intPointer()     // 8 字节
	float32Pointer() // 4 字节
	float64Pointer() // 8 字节
	boolPointer()    // 1 字节
	stringPointer()  // 16 字节
	StructPointer()  // 8 字节
}
