package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
)

// go run client.go
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	readerC := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(conn)

	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\r收到群消息: %s-> ", line)
		}
	}()

	go func() {
		for {
			fmt.Print("-> ")
			text, _ := readerC.ReadString('\n')
			conn.Write([]byte(text))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
