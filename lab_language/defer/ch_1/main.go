package main

import "fmt"

func test1() (x int) {
	defer fmt.Printf("in defer: x = %d\n", x) // 使用外部变量x=0
	x = 7
	return 9
}

func test2() (x int) { // 对比test3
	x = 7
	defer fmt.Printf("in defer: x = %d\n", x) // 使用外部变量x=7
	return 9
}

func test3() (x int) {
	x = 7
	defer func() {
		fmt.Printf("in defer: x = %d\n", x) // 使用外部变量x=9
	}()
	return 9
}

func test4() (x int) {
	defer func() {
		fmt.Printf("in defer: x = %d\n", x) // 使用外部变量x=9
	}()

	x = 7
	return 9
}

func test5() (x int) {
	defer func(n int) {
		fmt.Printf("in defer x as parameter: n = %d\n", n) // 使用局部变量n=0
		fmt.Printf("in defer x after return: x = %d\n", x) // 使用外部变量x=9
	}(x)

	x = 7
	return 9
}

func test6() (x int) {
	defer func(x int) {
		x = x + 5 // 使用局部变量x=0
	}(x)
	return 1 // test6=1
}

func test7() (x int) {
	defer func(x *int) {
		*x = *x + 5 // 使用局部指针变量x=0
	}(&x)
	return 1 // test7=6
}

func test8() (x int) {
	func() {
		x = x + 5 // 使用外部变量x=0
	}()
	return // test8=5
}

func test9() (x int) {
	x = 1
	func(x int) {
		x = x + 5 // 使用局部变量x=1
	}(x)
	return
}

// defer中传value类型，return后[无法改变改值]
// defer中传pointer类型，return后[可以改变改值]
// defer调用的是外部值，return后[可以改变改值]
func main() {
	fmt.Println("test1")
	fmt.Printf("in main: x = %d\n", test1())
	fmt.Println("test2")
	fmt.Printf("in main: x = %d\n", test2())
	fmt.Println("test3")
	fmt.Printf("in main: x = %d\n", test3())
	fmt.Println("test4")
	fmt.Printf("in main: x = %d\n", test4())
	fmt.Println("test5")
	fmt.Printf("in main: x = %d\n", test5())
	fmt.Println("test6")
	fmt.Printf("x = %d\n", test6())
	fmt.Println("test7")
	fmt.Printf("x = %d\n", test7())
	fmt.Println("test8")
	fmt.Printf("x = %d\n", test8())
	fmt.Println("test9")
	fmt.Printf("x = %d\n", test9())
}
