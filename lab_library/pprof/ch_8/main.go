package main

import (
	"net/http"
	"runtime"
)

func init() {
	println("init1 begin")
	slice := make([]int, 8)
	for i := 0; i < 32*1000*1000; i++ {
		slice = append(slice, i)
	}
}

// GODEBUG="inittrace=1" go run main.go
func main() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	select {}
}

/* GODEBUG GC: GODEBUG="inittrace=1" go run main.go
init main @3.1 ms, 2065 ms clock, 1293329824 bytes, 54 allocs

init main 			依赖包名
@3.1 ms 		 	程序启动到开始 init 中间消耗了 3.1 ms
2065 ms clock	    依赖包init动作所消耗了2065 ms(wall-clock)
1293329824 bytes    init动作在堆上分配的了 1293329824 bytes 内存
54 allocs		 	进行了 54 次 堆分配
*/
