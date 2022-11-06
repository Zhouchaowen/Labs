package main

import (
	"sync"
	"sync/atomic"
)

// Once 带返回的error
type Once struct {
	m sync.Mutex

	done uint32
}

// Do 传入函数f有返回值error，如果初始化失败，需要返回失败的error
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}

// 初始化
func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 { // 双检查，还没有初始化
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
