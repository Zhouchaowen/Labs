// https://www.yuque.com/qyuhen/go/tca6qu
package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 在同一底层数组的不同区间复制。
	src := s[5:8]
	dst := s[4:]

	n := copy(dst, src)
	fmt.Println(n, s)

	// 在不同数组间复制。
	dst = make([]int, 6)

	n = copy(dst, src)
	fmt.Println(n, dst)
}

/*
	3 [0 1 2 3 5 6 7 7 8 9]
	3 [6 7 7 0 0 0]
*/
