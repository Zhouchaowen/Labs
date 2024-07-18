package ch_2

import (
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
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

func BenchmarkMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idx := rand.Intn(N)
		_ = simpleMap[strconv.Itoa(idx)]
	}
}

func BenchmarkAtomic_Get(b *testing.B) {
	var ptr = unsafe.Pointer(&simpleMap)
	atomic.StorePointer(&ptr, unsafe.Pointer(&simpleMap))
	for i := 0; i < b.N; i++ {
		val := atomic.LoadPointer(&ptr)
		a := (*map[string]interface{})(val)
		idx := rand.Intn(N)
		_ = (*a)[strconv.Itoa(idx)]
	}
}

func BenchmarkMutexMap_Get(b *testing.B) {
	rlock := new(sync.RWMutex)
	for i := 0; i < b.N; i++ {
		rlock.RLock()
		idx := rand.Intn(N)
		_ = simpleMap[strconv.Itoa(idx)]
		rlock.RUnlock()
	}
}
