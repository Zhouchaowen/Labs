package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock uint32

func (sl *spinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched() //without this it locks up on GOMAXPROCS > 1
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

func (sl *spinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1)
}

func SpinLock() sync.Locker {
	var lock spinLock
	return &lock
}
