package main

import (
	"github.com/arl/statsviz"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// http://127.0.0.1:6060/debug/statsviz/
// 展示go程序运行时的指标
func main() {
	statsviz.RegisterDefault()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	<-time.After(10 * time.Second)
	fib(100)

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
}
