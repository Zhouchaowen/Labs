package main

import (
	"fmt"
	"time"
)

// chan传递指针,并改变指针地址的值
func main() {
	ch := make(chan []string, 0)
	slic := []string{"1", "2", "3"}

	go func() {
		ch <- slic
		slic[1] = "22"
	}()

	<-time.After(1 * time.Second)

	fmt.Println(<-ch)
}
