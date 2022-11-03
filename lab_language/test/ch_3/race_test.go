//  -race 参数，就可以在测试阶段发现可能的并发安全问题
package ch_3

import (
	"sync"
	"testing"
	"time"
)

// go test -run=TestParallelUnSafe -v -race .
func TestParallelUnSafe(t *testing.T) {
	a := 1
	go func() {
		a = 2
	}()
	a = 3
	t.Logf("a is %d", a)

	time.Sleep(2 * time.Second)
}

// go test -run=TestParallelSafe -v -race .
func TestParallelSafe(t *testing.T) {
	var mu sync.RWMutex
	a := 1
	go func() {
		mu.Lock()
		defer mu.Unlock()
		a = 2
	}()
	mu.Lock()
	a = 3
	t.Logf("a is %d", a)
	mu.Unlock()

	time.Sleep(2 * time.Second)
}
