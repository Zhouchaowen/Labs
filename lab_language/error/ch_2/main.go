package main

import (
	"errors"
	"fmt"
)

func Positive(n int) (bool, error) {
	if n == 0 {
		return false, errors.New("undefined")
	}
	return n > -1, nil
}

func Check(n int) {
	pos, err := Positive(n)
	if err != nil {
		fmt.Println("is neither")
		return
	}
	if pos {
		fmt.Println("is positive")
	} else {
		fmt.Println("is negative")
	}
}

/*
	异常和错误是两回事

	exceptions 破坏语义

	简单
	考虑失败，而不是成功（plan for failure, not success）
	没有隐藏的控制流
	完全交给你来控制 error
	Error are values

*/
func main() {
	Check(1)
	Check(0)
	Check(-1)
}
