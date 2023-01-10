package main

import "fmt"

func f() int {
	defer func() int {
		fmt.Println("ok")
		return 2
	}()

	return 1
}

func main() {
	fmt.Println(f())
}
