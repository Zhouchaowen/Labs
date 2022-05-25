package main

import "fmt"

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
}
