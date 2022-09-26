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

func gc() {
	for {
		_ = make([]byte, 16*Mi)
		time.Sleep(time.Second)
	}
}

// GODEBUG=gctrace=1 go run main.go | grep gc
// go tool pprof http://127.0.0.1:6060/debug/pprof/profile\?seconds\=20
func main() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	gc()
}

/* GODEBUG GC: GODEBUG=gctrace=1 go run main.go | grep gc
gc 1 @0.004s 2%: 0.017+0.47+0.002 ms clock, 0.017+0.27/0.13/0+0.002 ms cpu, 16->16->0 MB, 17 MB goal, 1 P
gc 2 @1.025s 0%: 0.020+0.41+0.003 ms clock, 0.020+0/0.061/0.32+0.003 ms cpu, 16->16->0 MB, 17 MB goal, 1 P
gc 3 @2.031s 0%: 0.040+0.58+0.003 ms clock, 0.040+0/0.11/0.42+0.003 ms cpu, 16->16->0 MB, 17 MB goal, 1 P
gc 4 @3.036s 0%: 0.041+0.57+0.004 ms clock, 0.041+0/0.12/0.41+0.004 ms cpu, 16->16->0 MB, 17 MB goal, 1 P

gc 1        					第1次执行GC
@0.004s     					程序已经执行了0.004秒(我们可以看到这列数据一直在递增)
2%          					gc时间占程序总执行时间的2%
0.017+0.47+0.002 ms clock 		垃圾回收各阶段占用的时间(wall-clock，现实意义消耗的时间): STW(stop-the-world)清扫终止+并发标记和扫描的时间+STW标记终止的时间。
0.017+0.27/0.13/0+0.002 ms cpu  也是gc各阶段占用的时间，但是程序在cpu上消耗的时间。 STW(stop-the-world)清扫的时间+并发标记和扫描的时间(辅助时间/后台gc时间/闲置gc时间)+STW标记的时间。
16->16->0 MB	        		堆在gc开始时的大小、gc结束时的大小、当前活跃的大小
17 MB goal						全局堆的大小
1 P								P(process)的数量是1

*/
