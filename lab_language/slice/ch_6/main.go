// https://www.yuque.com/qyuhen/go/cmc1bu
package main

import "fmt"

// FILO: 先进后出，栈。

type Stack []int

func NewStack() *Stack {
	s := make(Stack, 0, 10)
	return &s
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	x, n := *s, len(*s)

	v := x[n-1]
	*s = x[:n-1]

	return v, true
}

// ---------------------------

func main() {
	s := NewStack()

	// push
	for i := 0; i < 5; i++ {
		s.Push(i + 10)
	}

	// pop
	for i := 0; i < 7; i++ {
		fmt.Println(s.Pop())
	}
}

/*

14 true
13 true
12 true
11 true
10 true
0 false
0 false

*/
