// https://github.com/kevinyan815/gocookbook/issues/22
package main

import (
"fmt"
"math/rand"
"strings"
"sync"
)

type Queue struct {
	cond *sync.Cond
	data []interface{}
	capc int
	logs []string
}

func NewQueue(capacity int) *Queue {
	return &Queue{cond: &sync.Cond{L: &sync.Mutex{}}, data: make([]interface{}, 0), capc: capacity, logs: make([]string, 0)}
}

func (q *Queue) Enqueue(d interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.data) == q.capc {
		q.cond.Wait()
	}
	// FIFO入队
	q.data = append(q.data, d)
	// 记录操作日志
	q.logs = append(q.logs, fmt.Sprintf("En %v\n", d))
	// 通知其他waiter进行Dequeue或Enqueue操作
	q.cond.Broadcast()

}

func (q *Queue) Dequeue() (d interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.data) == 0 {
		q.cond.Wait()
	}
	// FIFO出队
	d = q.data[0]
	q.data = q.data[1:]
	// 记录操作日志
	q.logs = append(q.logs, fmt.Sprintf("De %v\n", d))
	// 通知其他waiter进行Dequeue或Enqueue操作
	q.cond.Broadcast()
	return
}

func (q *Queue) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.data)
}

func (q *Queue) String() string {
	var b strings.Builder
	for _, log := range q.logs {
		//fmt.Fprint(&b, log)
		b.WriteString(log)
	}
	return b.String()
}


func Example() {
	var wg sync.WaitGroup
	//容量为5的阻塞队列
	que := NewQueue(3)

	// 生成随机命令
	for i, cmd := range Commands(20, true) {
		wg.Add(1)

		// 0表示入队，1表示出队
		if cmd == 0 {
			go func(id int) {
				defer wg.Done()
				que.Enqueue(id)
			}(i)
		} else {
			go func(id int) {
				defer wg.Done()
				que.Dequeue()
			}(i)
		}
	}

	/*
		// 当执行出队、入队命令的worker数『不相等』时
		// 最后会有worker阻塞在出队或入队方法上
		// 同时主goroutine会阻塞在wg.Wait()上
		// 此时所有goroutine都阻塞了
		// 下面的goroutine会避免该问题
		// 但仍需新worker唤醒阻塞在队列上的worker
		go func() {
			for {
				select{
				case <-time.After(time.Second):
					runtime.Gosched()
				}
			}
		}()
	*/

	wg.Wait()

	// 输出操作日志
	fmt.Println(que)
}

// Commands 用于产生出队、入队命令
func Commands(N int, random bool) []int {
	if N%2 != 0 {
		panic("will deadlock!")
	}
	// 0表示入队，1表示出队
	commands := make([]int, N)
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			commands[i] = 1
		}
	}

	if random {
		// shuffle algorithms
		for i := len(commands) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			commands[i], commands[j] = commands[j], commands[i]
		}
	}

	return commands
}

