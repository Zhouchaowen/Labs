package main

import "fmt"

/*
调用者要使用类型断言和类型 switch，就要让自定义的 error 变为 public。这种模型会导致和调用者产生强耦合，从而导致 API 变得脆弱。

结论是尽量避免使用 error types，虽然错误类型比 sentinel errors 更好，因为它们可以捕获关于出错的更多上下文，但是 error types 共享 error values 许多相同的问题。

因此，我的建议是避免错误类型，或者至少避免将它们作为公共 API 的一部分。
*/
type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"something happend", "saver.go", 23}
}

// 与错误值相比，错误类型的一大改进是它们能够包装底层错误以提供更多上下文
func main() {
	err := test()
	switch err := err.(type) {
	case nil:
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:
	}
}

/*
// 使用案例
// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op   string
	Path string
	Err  error
}
*/
