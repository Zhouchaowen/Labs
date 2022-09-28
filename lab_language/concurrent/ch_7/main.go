package main

import "fmt"
import "time"

// 这里是 worker，我们将并发执行多个 worker。
// worker 将从 `jobs` 通道接收任务，并且通过 `results` 发送对应的结果。
// 我们将让每个任务间隔 1s 来模仿一个耗时的任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
	close(results)
}

func main() {

	// 为了使用 worker 线程池并且收集他们的结果，我们需要 2 个通道。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 这里我们发送 9 个 `jobs`，然后 `close` 这些通道
	// 来表示这些就是所有的任务了。
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后，我们收集所有这些任务的返回值。
	for v := range results {
		fmt.Println(v)
	}
}
