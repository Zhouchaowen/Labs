// 共享底层存储导致扩容时的坑
package main

import "fmt"

func dome1() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	fmt.Printf("dome_1 start  : %p\n", s)
	modifySlice1(s)
	fmt.Printf("dome_1 end    : %p\n", s)
	fmt.Println("dome_1 ", s) //[1024,1,2]
}

func modifySlice1(s []int) {
	fmt.Printf("dome_1 center1: %p\n", s)
	s[0] = 1024
	fmt.Printf("dome_1 center2: %p\n", s)
}

// 对比dome3
func dome2() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	fmt.Printf("dome_2 start  : %p\n", s)
	modifySlice2(s)
	fmt.Printf("dome_2 end    : %p\n", s)
	fmt.Println("dome_2 ", s) //[1024,1,2]
}

func modifySlice2(s []int) {
	fmt.Printf("dome_2 center1: %p\n", s)
	s = append(s, 2048)
	s[0] = 1024
	fmt.Printf("dome_2 center2: %p\n", s)
}

// 对比dome4
func dome3() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	fmt.Printf("dome_3 start  : %p\n", s)
	modifySlice3(s)
	fmt.Printf("dome_3 end    : %p\n", s)
	fmt.Println("dome_3 ", s) //[0,1,2]
}

func modifySlice3(s []int) {
	fmt.Printf("dome_3 center1: %p\n", s)
	s = append(s, 2048)
	s = append(s, 4096)
	s[0] = 1024
	fmt.Printf("dome_3 center2: %p\n", s)
}

func dome4() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	fmt.Printf("dome_4 start  : %p\n", s)
	modifySlice4(s)
	fmt.Printf("dome_4 end    : %p\n", s)
	fmt.Println("dome_4 ", s) //[1024,1,2]
}

func modifySlice4(s []int) {
	fmt.Printf("dome_4 center1: %p\n", s)
	s[0] = 1024
	s = append(s, 2048)
	s = append(s, 4096)
	fmt.Printf("dome_4 center2: %p\n", s)
}

func main() {
	dome1()
	fmt.Println("----------------")
	dome2()
	fmt.Println("----------------")
	dome3()
	fmt.Println("----------------")
	dome4()
}
