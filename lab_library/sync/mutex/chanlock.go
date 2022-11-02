package main

type ChanMutex chan struct{}

func (m *ChanMutex) Lock() {
	ch := (chan struct{})(*m)
	ch <- struct{}{}
}

func (m *ChanMutex) Unlock() {
	ch := (chan struct{})(*m)
	select {
	case <-ch:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (m *ChanMutex) TryLock() bool {
	ch := (chan struct{})(*m)
	select {
	case ch <- struct{}{}:
		return true
	default:
	}
	return false
}
