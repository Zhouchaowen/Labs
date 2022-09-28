package ch_4

import (
	"sync"
)

type MutexMap struct {
	items map[string]string
	mu    sync.RWMutex
}

func NewMutexMap() MutexMap {
	return MutexMap{
		items: map[string]string{},
	}
}

func (m *MutexMap) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.items[key]
	return value, ok
}

func (m *MutexMap) Set(key string, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[key] = value
}
