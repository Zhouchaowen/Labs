// 实验，不加锁替换map导致数据混乱
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type A struct {
	Mp map[string]int
}

func UnLockMap() {
	var a A
	arrMap := make([]map[string]int, 5)
	for i := 0; i < 5; i++ {
		arrMap[i] = make(map[string]int)
	}

	ctx, cannel := context.WithCancel(context.Background())

	go func() {
		a = A{arrMap[0]}
		for {
			select {
			case <-ctx.Done():
				return
			default:
				a.Mp["1"] += 1
			}
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println(arrMap[0])
		a = A{arrMap[1]}
	}()

	time.Sleep(7 * time.Second)
	cannel()
	fmt.Printf("%v\n", arrMap)
}

func LockMap() {
	var a A
	arrMap := make([]map[string]int, 5)
	for i := 0; i < 5; i++ {
		arrMap[i] = make(map[string]int)
	}
	var mu sync.RWMutex
	ctx, cannel := context.WithCancel(context.Background())

	go func() {
		a = A{arrMap[0]}
		for {
			select {
			case <-ctx.Done():
				return
			default:
				mu.Lock()
				a.Mp["1"] += 1
				mu.Unlock()
			}
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		mu.Lock()
		fmt.Println(arrMap[0])
		a = A{arrMap[1]}
		mu.Unlock()
	}()

	time.Sleep(7 * time.Second)
	cannel()
	fmt.Printf("%v\n", arrMap)
}

func main() {
	UnLockMap()
	LockMap()
}
