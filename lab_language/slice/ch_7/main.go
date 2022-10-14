// https://www.yuque.com/qyuhen/go/cmc1bu
package main

import "fmt"

// FIFO：先进先出，队列。

type Queue []int

func NewQueue() *Queue {
	q := make(Queue, 0, 10)
	return &q
}

func (q *Queue) Put(v int) {
	*q = append(*q, v)
}

func (q *Queue) Get() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}

	x := *q
	v := x[0]

	// copy(x, x[1:])
	// *q = x[:len(x) - 1]

	*q = append(x[:0], x[1:]...) // 等同上两行。

	return v, true
}

// ---------------------------

func main() {
	q := NewQueue()

	// put
	for i := 0; i < 5; i++ {
		q.Put(i + 10)
	}

	// get
	for i := 0; i < 7; i++ {
		fmt.Println(q.Get())
	}
}

/*

10 true
11 true
12 true
13 true
14 true
0 false
0 false

*/
