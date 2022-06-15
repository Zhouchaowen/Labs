package main

import (
	"fmt"
)

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

func addSlice() {
	a := []int{1, 2, 3, 4, 5, 6}
	t := len(a)
	aMap := make(map[int]bool)

	for i := 0; i < t; i++ {
		if aMap[a[i]] {
			continue
		}
		aMap[a[i]] = true
		if a[i] == 3 || a[i] == 9 {
			a = append(a, []int{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}...)
			t += 20
		}

		fmt.Println(a[i])
	}
}

func main() {
	//modifySlice()
	addSlice()
}
