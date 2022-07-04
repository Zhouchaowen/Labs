package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var HOST = "127.0.0.1:8080"

func server() {
	listener, err := net.Listen("tcp", HOST)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", err)
			return
		}

		go func(conn net.Conn) {
			w := bufio.NewWriter(conn)
			r := bufio.NewReader(conn)
			for {
				ping, err := r.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("server receive msg -> %s", ping)

				_, err = w.WriteString("pong\n")
				if err != nil {
					log.Fatal(err)
				}
				w.Flush()
				fmt.Println("server send msg -> pong")
			}
		}(conn)
	}
}

func main() {
	server()
}
