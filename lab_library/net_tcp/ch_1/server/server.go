package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var HOST = "127.0.0.1:8080"

func server() {
	//1、监听端口
	listener, err := net.Listen("tcp", HOST)
	if err != nil {
		log.Fatal(err)
	}
	//2.建立套接字连接
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

		//3. 创建处理协程
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
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
}

func main() {
	server()
}
