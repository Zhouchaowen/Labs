package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

// go run -gcflags '-m -l' main.go
func main() {
	x := foo()
	fmt.Println(*x)
}
