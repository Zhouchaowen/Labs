// 只操作一次
// 1.初始化单例资源
// 2.并发访问只需要初始化一次的共享资源
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var connMu sync.Mutex
var conn net.Conn

// Single 单例方式一: 每次请求会存在锁竞争
func Single() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	if conn != nil {
		return conn
	}

	conn, _ := net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}

func example() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody) // 只有第一次调用才执行
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// Output:
	// Only once
}

func main() {
	example()
}
