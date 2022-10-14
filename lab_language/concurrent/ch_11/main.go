package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}
func (m *Mutex) Lock() {
	<-m.ch
}
func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}

func main() {
	mp := make(map[int]int)
	mx := NewMutex()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Panic(err)
			}
		}()
		for {
			select {
			case <-ctx.Done():
				log.Println("read done")
				return
			default:
				mx.Lock()
				mp[rand.Intn(2)] += 1
				mx.Unlock()
			}
			time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Panic(err)
			}
		}()
		for {
			select {
			case <-ctx.Done():
				log.Println("read done")
				return
			default:
				if mx.TryLock() {
					fmt.Println("read:", mp[rand.Intn(2)])
					mx.Unlock()
				}
			}
			time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		}
	}()

	<-time.After(15 * time.Second)
}
