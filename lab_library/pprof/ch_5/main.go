package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

const (
	Ki = 1024
	Mi = Ki * Ki
	Gi = Ki * Mi
	Ti = Ki * Gi
	Pi = Ki * Ti
)

func coroutinesLeaked() {
	for {
		for i := 0; i < 10; i++ {
			go func() {
				time.Sleep(30 * time.Second)
			}()
		}
		time.Sleep(time.Second)
	}
}

// go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine\?seconds\=20
func main() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	coroutinesLeaked()
}
