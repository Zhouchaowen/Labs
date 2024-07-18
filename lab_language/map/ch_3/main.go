package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"unsafe"
)

var N = 10000

var simpleMap map[string]interface{}

func init() {
	simpleMap = make(map[string]interface{}, N)
	for j := 0; j < N; j++ {
		simpleMap[strconv.Itoa(j)] = j + 1
	}
}

func main() {
	var ptr = unsafe.Pointer(&simpleMap)
	atomic.StorePointer(&ptr, unsafe.Pointer(&simpleMap))
	val := atomic.LoadPointer(&ptr)
	a := (*map[string]interface{})(val)
	fmt.Println((*a)["0"])
}
