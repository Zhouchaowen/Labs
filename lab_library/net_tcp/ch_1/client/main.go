package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

var HOST = "127.0.0.1:8080"

func client() {
	conn, err := net.Dial("tcp", HOST)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(conn)
	r := bufio.NewReader(conn)

	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				// 1、回复服务器信息
				_, err := w.WriteString("ping\n")
				if err != nil {
					log.Fatal(err)
				}
				w.Flush()

				fmt.Println("client send msg -> ping")
			}
		}
	}()

	go func() {
		for {
			// 1、一直读取直到读到\n
			pong, err := r.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("client receive msg -> %s", pong)
		}
	}()

	select {
	case <-time.After(1 * time.Minute):
	}

}

func main() {
	client()
}
