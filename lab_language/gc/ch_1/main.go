// 基础类型slice Gc问题

// int,bool,float 等基础类型组成的数组 GC扫描时会忽略
package main

import (
	"fmt"
	"runtime"
	"time"
)

// 指针类型数组，GC扫描耗时 GC took 14.190040366s
func func1() {
	a := make([]*int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

// int类型数组，GC扫描耗时 GC took 4.454137ms
func func2() {
	a := make([]int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

// int类型二维数组，GC扫描耗时 GC took GC took 43.88830211s
func func3() {
	a := make([][]int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

// interface类型数组，GC扫描耗时 GC took 39.226903503s
func func4() {
	a := make([]interface{}, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

func func5() {
	a := make([]*string, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

func func6() {
	a := make([]string, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

// GODEBUG=gctrace=1 go run main.go | grep gc
func main() {
	//func1()
	//func2()
	//func3()
	//func4()
	//func5()
	func6()
}
