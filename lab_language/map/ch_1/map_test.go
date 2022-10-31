// https://shaffer.cn/golang/golang-map-benchmark/
// go test -bench=. -benchtime=10s -benchmem
// pkg: Labs/lab_language/map/ch_1
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkMap_Set-4                  4743           2169202 ns/op          156164 B/op      19646 allocs/op
// BenchmarkMap_Get-4                  7833           1718128 ns/op           39614 B/op       9901 allocs/op
// BenchmarkMutexMap_Set-4             1825           6879003 ns/op          406854 B/op      29857 allocs/op
// BenchmarkMutexMap_Get-4             2318           4657277 ns/op          202349 B/op      19929 allocs/op
// BenchmarkSyncMap_Set-4              1224          10245496 ns/op          734348 B/op      49974 allocs/op
// BenchmarkSyncMap_Get-4              2478           4548111 ns/op          201663 B/op      19921 allocs/op
// BenchmarkConcurrentMap_Set-4        2370           5500791 ns/op          399756 B/op      29755 allocs/op
// BenchmarkConcurrentMap_Get-4        2516           4684854 ns/op          200949 B/op      19915 allocs/op
package main

import (
	"github.com/orcaman/concurrent-map"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

var N = 10000

var simpleMap map[string]interface{}
var mutexMap MutexMap
var syncMap sync.Map
var concurrentMap cmap.ConcurrentMap

func init() {
	simpleMap = make(map[string]interface{}, N)
	mutexMap = MutexMap{m: make(map[string]interface{}, N)}
	syncMap = sync.Map{}
	for i := 0; i < N; i++ {
		syncMap.Store(strconv.Itoa(i), i)
	}
	concurrentMap = cmap.New()
	for i := 0; i < N; i++ {
		concurrentMap.Set(strconv.Itoa(i), i)
	}
}

func BenchmarkMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			idx := rand.Intn(N)
			simpleMap[strconv.Itoa(idx)] = j + 1
			wg.Done()
		}
		wg.Wait()
	}
}

func BenchmarkMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			idx := rand.Intn(N)
			_ = simpleMap[strconv.Itoa(idx)]
			wg.Done()
		}
		wg.Wait()
	}
}

func BenchmarkMutexMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				mutexMap.Set(strconv.Itoa(idx), j+1)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkMutexMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				mutexMap.Get(strconv.Itoa(idx))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				syncMap.Store(strconv.Itoa(idx), j+1)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				syncMap.Load(strconv.Itoa(idx))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				concurrentMap.Set(strconv.Itoa(idx), j+1)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				concurrentMap.Get(strconv.Itoa(idx))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

type MutexMap struct {
	sync.RWMutex
	m map[string]interface{}
}

func (m *MutexMap) Set(k string, v interface{}) {
	m.Lock()
	m.m[k] = v
	m.Unlock()
}

func (m *MutexMap) Get(k string) (v interface{}) {
	m.RLock()
	v = m.m[k]
	m.RUnlock()
	return
}
