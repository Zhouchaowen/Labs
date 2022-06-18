package ch_4

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

//https://www.cnblogs.com/yinbiao/p/15884420.html

/*
https://shaffer.cn/golang/golang-map-benchmark/
https://segmentfault.com/a/1190000018448064
https://blog.csdn.net/wyg_031113/article/details/106282340
https://xiaorui.cc/archives/5611
https://github.com/halfrost/Halfrost-Field/blob/master/contents/Go/go_map_chapter_one.md

*/

const (
	count           = 1e6
	concurrentCount = 10
)

// go test -gcflags "-N -l" -bench='^Benchmark' -benchmem .
func BenchmarkShardMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		testKV := map[string]string{}
		for i := 0; i < count; i++ {
			testKV["key_"+strconv.Itoa(i)] = "value_" + strconv.Itoa(i)
		}

		sm := NewMapShards()

		for k, v := range testKV {
			sm.Set(k, v)
		}

		wg := sync.WaitGroup{}
		wg.Add(concurrentCount * 2)

		b.StartTimer()

		for l := 0; l < concurrentCount; l++ {
			go func() {
				defer wg.Done()
				for i := 0; i < count; i++ {
					index := rand.Intn(1e9)
					sm.Set("key_"+strconv.Itoa(index), "value_"+strconv.Itoa(index))
				}
			}()
		}

		for l := 0; l < concurrentCount; l++ {
			go func() {
				defer wg.Done()
				for i := 0; i < count; i++ {
					index := rand.Intn(1e9)
					sm.Get("key_" + strconv.Itoa(index))
				}
			}()
		}

		wg.Wait()
	}
}

func BenchmarkMutexMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		testKV := map[string]string{}
		for i := 0; i < count; i++ {
			testKV["key_"+strconv.Itoa(i)] = "value_" + strconv.Itoa(i)
		}

		mm := NewMutexMap()

		for k, v := range testKV {
			mm.Set(k, v)
		}

		wg := sync.WaitGroup{}
		wg.Add(concurrentCount * 2)

		b.StartTimer()

		for l := 0; l < concurrentCount; l++ {
			go func() {
				defer wg.Done()
				for i := 0; i < count; i++ {
					index := rand.Intn(1e9)
					mm.Set("key_"+strconv.Itoa(index), "value_"+strconv.Itoa(index))
				}
			}()
		}

		for l := 0; l < concurrentCount; l++ {
			go func() {
				defer wg.Done()
				for i := 0; i < count; i++ {
					index := rand.Intn(1e9)
					mm.Get("key_" + strconv.Itoa(index))
				}
			}()
		}

		wg.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		testKV := map[string]string{}
		for i := 0; i < count; i++ {
			testKV["key_"+strconv.Itoa(i)] = "value_" + strconv.Itoa(i)
		}

		sm := &sync.Map{}

		for k, v := range testKV {
			sm.Store(k, v)
		}

		wg := sync.WaitGroup{}
		wg.Add(concurrentCount * 2)

		b.StartTimer()

		for l := 0; l < concurrentCount; l++ {
			go func() {
				defer wg.Done()
				for i := 0; i < count; i++ {
					index := rand.Intn(1e9)
					sm.Store("key_"+strconv.Itoa(index), "value_"+strconv.Itoa(index))
				}
			}()
		}

		for l := 0; l < concurrentCount; l++ {
			go func() {
				defer wg.Done()
				for i := 0; i < count; i++ {
					index := rand.Intn(1e9)
					sm.Load("key_" + strconv.Itoa(index))
				}
			}()
		}

		wg.Wait()
	}
}

// go test -v -run ^Test
func TestMutexMap(t *testing.T) {
	sm := NewMutexMap()
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			sm.Set("key_"+strconv.Itoa(i), "value_"+strconv.Itoa(i))
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			if value, ok := sm.Get("key_" + strconv.Itoa(i)); ok {
				fmt.Println(value)
			}
		}
	}()
	wg.Wait()
}
