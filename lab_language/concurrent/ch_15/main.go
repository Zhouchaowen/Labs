package main

import (
	"fmt"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type Test struct {
	X int
	Y int
}

func NoMutexConcurrent1() {
	var g Test

	for i := 0; i < 1000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				g = Test{1, 2}
			}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				g = Test{3, 4}
			}
		}()

		// 协程 3
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				// 赋值异常判断
				if !((g.X == 1 && g.Y == 2) || (g.X == 3 && g.Y == 4)) {
					fmt.Printf("concurrent assignment error, i=%v g=%+v\n", i, g)
					break
				}
			}
		}()
		wg.Wait()
	}
}

func NoMutexConcurrent2() {
	var g *Test
	g = &Test{1, 2}

	for i := 0; i < 1000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				g = &Test{1, 2}
			}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				g = &Test{3, 4}
			}
		}()

		// 协程 3
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				// 赋值异常判断
				if !((g.X == 1 && g.Y == 2) || (g.X == 3 && g.Y == 4)) {
					fmt.Printf("concurrent assignment error, i=%v g=%+v", i, g)
					break
				}
			}
		}()
		wg.Wait()
	}
}

func NoMutexConcurrent3() {
	g := make(map[string]string)

	for i := 0; i < 100; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				tmp := map[string]string{
					"1": "10",
					"2": "20",
				}
				g = tmp
			}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				tmp := map[string]string{
					"3": "30",
					"4": "40",
					"5": "50",
				}
				g = tmp
			}
		}()

		// 协程 3
		wg.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("panic error: %s，stack=%s\n", err, string(debug.Stack()))
				}
				wg.Done()

			}()
			for i := 0; i < 10000; i++ {
				// 赋值异常判断
				if !((g["1"] == "10" && g["2"] == "20") || (g["3"] == "30" && g["4"] == "40")) {
					fmt.Printf("concurrent assignment error, i=%v g=%+v\n", i, g)
					break
				}
			}
		}()
		wg.Wait()
	}
}

func NoMutexConcurrent4() {
	var g *map[int]int
	tmp := map[int]int{
		1: 10,
		2: 20,
	}
	g = &tmp

	for i := 0; i < 100000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				tmp1 := map[int]int{
					1: 10,
					2: 20,
				}
				g = &tmp1
				time.Sleep(1 * time.Millisecond)
			}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				tmp2 := map[int]int{
					3: 30,
					4: 40,
					5: 50,
				}
				g = &tmp2
				time.Sleep(1 * time.Millisecond)
			}
		}()

		// 协程 3
		wg.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("panic error: %s，stack=%s\n", err, string(debug.Stack()))
				}
				wg.Done()

			}()
			for i := 0; i < 10000; i++ {
				// 赋值异常判断
				//a := *g
				// 特别注意 必须要a := *g先取出来 否则如下将会报错
				if !(((*g)[1] == 10 && (*g)[2] == 20) || ((*g)[3] == 30 && (*g)[4] == 40)) {
					fmt.Printf("concurrent assignment error, i=%v g=%+v\n", i, *g)
					break
				}
				//fmt.Println(a[1], "+++", a[2], "+++", a[3], "+++", a[4])
				//if !((a[1] == 10 && a[2] == 20) || (a[3] == 30 && a[4] == 40)) {
				//	fmt.Printf("concurrent assignment error, i=%v g=%+v\n", i, *g)
				//	break
				//}
				time.Sleep(1 * time.Millisecond)
			}
		}()
		wg.Wait()
	}
}

func NoMutexConcurrent5() {
	var g atomic.Value
	g.Store(map[string]string{})

	for i := 0; i < 100000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				g.Swap(map[string]string{
					"1": "10",
					"2": "20",
				})
			}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				g.Swap(map[string]string{
					"3": "30",
					"4": "40",
					"5": "50",
				})
			}
		}()

		// 协程 3
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				// 赋值异常判断
				mp := g.Load().(map[string]string)
				if !((mp["1"] == "10" && mp["2"] == "20") || (mp["3"] == "30" && mp["4"] == "40")) {
					fmt.Printf("concurrent assignment error, i=%v g=%+v\n", i, g)
					break
				}
			}
		}()
		wg.Wait()
	}
}

// 并发赋值 复合值的安全问题
func main() {
	//NoMutexConcurrent1()
	//NoMutexConcurrent2()
	NoMutexConcurrent3()
	//NoMutexConcurrent4()
	//NoMutexConcurrent5()
}

/*
	https://github.com/Terry-Mao/gopush-cluster/issues/44
	https://wnanbei.github.io/post/go-%E5%8E%9F%E5%AD%90%E6%93%8D%E4%BD%9C-atomic/
	https://learnku.com/go/t/76112
	https://cloud.tencent.com/developer/beta/article/1810536
	https://gfw.go101.org/article/concurrent-atomic-operation.html
*/
