// 结构体类型 Gc问题

// 结构体只包含基础类型，组成的数组GC扫描时会忽略它
package main

import (
	"fmt"
	"runtime"
	"time"
)

type t struct {
	p, q int32
}

// a存储的数据时指针，gc会扫描数组每个原始，检查如果是指针会继续扫描下去
func func1() {
	a := make([]*t, 1e9)
	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

// gc耗时 GC took 655.397µs
func func2() {
	a := make([]t, 1e9)
	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

// t3.str string 是引用类型 GC会扫描
type t3 struct {
	p   int32
	str string
}

func func3() {
	a := make([]t3, 1e9)
	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

// t4.str []byte 是引用类型 GC会扫描
type t4 struct {
	p   int32
	str []byte
}

func func4() {
	a := make([]t4, 1e9)
	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

// t5.t time.Time 是引用类型 GC会扫描
type t5 struct {
	p int32
	t time.Time
}

func func5() {
	a := make([]t5, 1e9)
	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

type t6 struct {
	p int32
	f bool
}

func func6() {
	a := make([]t6, 1e9)
	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

type t7 struct {
	p interface{}
}

func func7() {
	a := make([]t7, 1e9)
	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}

func main() {
	//func1()
	//func2()
	//func3()
	//func4()
	//func5()
	func6()
	//func7()
}
