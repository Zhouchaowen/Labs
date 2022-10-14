// https://www.yuque.com/qyuhen/go/cmc1bu
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	sync.Mutex
	data []int
	head int
	tail int
}

func NewQueue(cap int) *Queue {
	return &Queue{data: make([]int, cap)}
}

func (q *Queue) Put(v int) bool {
	q.Lock()
	defer q.Unlock()

	if q.tail-q.head == len(q.data) {
		return false
	}

	q.data[q.tail%len(q.data)] = v
	q.tail++

	return true
}

func (q *Queue) Get() (int, bool) {
	q.Lock()
	defer q.Unlock()

	if q.tail-q.head == 0 {
		return 0, false
	}

	v := q.data[q.head%len(q.data)]
	q.head++

	return v, true
}

// ---------------------------

func main() {
	rand.Seed(time.Now().UnixNano())

	const max = 100000
	src := rand.Perm(max) // 随机测试数据。
	dst := make([]int, 0, max)

	q := NewQueue(6)

	// ------------------------

	var wg sync.WaitGroup
	wg.Add(2)

	// put
	go func() {
		defer wg.Done()
		for _, v := range src {
			for {
				if ok := q.Put(v); !ok {
					continue
				}
				break
			}
		}
	}()

	// get
	go func() {
		defer wg.Done()
		for len(dst) < max {
			if v, ok := q.Get(); ok {
				dst = append(dst, v)
				continue
			}
		}
	}()

	wg.Wait()

	// 转换成数组进行比较。
	if *(*[max]int)(src) != *(*[max]int)(dst) {
		log.Fatalln("fail !!!")
	}

	log.Printf("%+v\n", *q)
}

// {data:[99011 52214 53425 10572 82360 78821] head:100000 tail:100000}
