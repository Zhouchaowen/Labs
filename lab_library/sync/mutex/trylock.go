// https://colobu.com/2017/03/09/implement-TryLock-in-Go/
package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	// 尝试抢锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果处于唤醒，加锁，饥饿状态，这次请求就不参与竞争了，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// 尝试在竞争状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, new)
}
