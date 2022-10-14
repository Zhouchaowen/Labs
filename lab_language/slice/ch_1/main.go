// slice 切片被修改影响底层数据
// https://www.yuque.com/qyuhen/go/tca6qu
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	# 切片
	+---+---+---+---+---+---+---+---+---+---+
	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |   array
	+---+---+---+---+---+---+---+---+---+---+
			|               |       |           slice: [low : high : max]
			|<--- s.len --->|       |
			|                       |           len = high - low
			|<------- s.cap ------->|           cap = max  - low
*/
/*
	# 切片引用原数组
	+---+---+---+---+---+---+---+---+---+---+
	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |   a: [10]int
	+---+---+---+---+---+---+---+---+---+---+
			.                       .
			+---+---+---+---+---+---+
			| 2 | 3 | 4 | 5 |   |   |           s: a[2:6:8]
			+---+---+---+---+---+---+
			0   1   2   3

	s[0]: 2
	s[1]: 3
	s[2]: 4
	s[3]: 5
*/
/*
	# 切片的数据来自底层数组，重切片只是调整引用范围。
	# 重切片受原 cap 限制，而非 len。

			+---+---+---+---+---+
			| 3 | 4 | 5 | 6 | 7 |     s2: s[1:6]
			+---+---+---+---+---+
			.                   .
		+---+---+---+---+---+---+
		| 2 | 3 | 4 | 5 |   |   |     s
		+---+---+---+---+---+---+
		.               .
		+---+---+---+---+
		| 2 | 3 |   |   |             s1: s[0:2:4]
		+---+---+---+---+

	s1: [2 3],       len = 2, cap = 4
	s2: [3 4 5 6 7], len = 5, cap = 5
*/

func p(s []int) {
	fmt.Printf("%t, %d, %#v\n",
		s == nil,
		unsafe.Sizeof(s),
		(*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func memory() {

	// 仅分配 header 内存，未初始化。
	var s1 []int

	// 初始化。
	s2 := []int{}

	// 调用 makeslice 初始化。
	s3 := make([]int, 0)

	p(s1)
	p(s2)
	p(s3)
}

func modifySlice() {
	var s = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	b := s[0:5]

	fmt.Println("s:", s)
	fmt.Println("b:", b)
	fmt.Println("&s[0]:", &s[0])
	fmt.Println("&b[0]:", &b[0])
	fmt.Printf("&s:%p\n", s)
	fmt.Printf("&b:%p\n", b)
	b[0] = 'A'
	fmt.Println("=========================")
	fmt.Println("s:", s)
	fmt.Println("b:", b)
	fmt.Println("&s[0]:", &s[0])
	fmt.Println("&b[0]:", &b[0])
	fmt.Printf("&s:%p\n", s)
	fmt.Printf("&b:%p\n", b)
}

func main() {
	modifySlice()
	fmt.Println("=========================")
	memory()
}
