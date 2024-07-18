package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Dome struct {
	a string
	b string
}

type DomeA struct {
	c []*Dome
}

var Domes []*Dome

func init() {
	Domes = []*Dome{{
		a: ".",
		b: "enum.EcsPolicyLoadOriginalWc",
	}}
}

func main() {
	var wg sync.WaitGroup
	fmt.Printf("%+v\n", Domes[0])

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			a := DomeA{
				c: Domes,
			}
			fmt.Printf("%+v\n", a.c[0])
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			runtime.GC()
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	runtime.GC()
	time.Sleep(3 * time.Second)
	fmt.Printf("%+v\n", Domes[0])
}
