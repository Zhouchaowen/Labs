package main

import "fmt"

func dome1() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice1(s)
	fmt.Println("dome_1 ", s) //[1024,1,2]
}

func modifySlice1(s []int) {
	s[0] = 1024
}

func dome2() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice2(s)
	fmt.Println("dome_2 ", s) //[1024,1,2]
}

func modifySlice2(s []int) {
	s = append(s, 2048)
	s[0] = 1024
}

func dome3() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice3(s)
	fmt.Println("dome_3 ", s) //[1024,1,2]
}

func modifySlice3(s []int) {
	s = append(s, 2048)
	s = append(s, 4096)
	s[0] = 1024
}

func dome4() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice4(s)
	fmt.Println("dome_4 ", s) //[1024,1,2]
}

func modifySlice4(s []int) {
	s[0] = 1024
	s = append(s, 2048)
	s = append(s, 4096)
}

func main() {
	dome1()
	dome2()
	dome3()
	dome4()
}
