package main

func main() {
	// 创建一个int类型的通道
	ch := make(chan int)

	// 开启一个匿名 goroutine
	go func() {
		// 向通道发送数字42
		ch <- 42
	}()
	// 从通道中读取
	<-ch
}
