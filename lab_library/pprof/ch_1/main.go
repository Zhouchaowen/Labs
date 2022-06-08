package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	if n <= 1 {
		return 1
	}

	f1, f2 := 1, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}

// go tool pprof http://10.2.8.17:6060/debug/pprof/profile\?seconds\=20
func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	for {
		time.Sleep(5 * time.Second)
		n := 10
		for i := 1; i <= 5; i++ {
			fmt.Printf("fib(%d)=%d\n", n, fib(n))
			n += 3 * i
		}
	}
}
