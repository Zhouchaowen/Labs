package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}

var once sync.Once

func getInstanceOnce() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creting Single Instance Now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		//go getInstance()
		go getInstanceOnce()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
