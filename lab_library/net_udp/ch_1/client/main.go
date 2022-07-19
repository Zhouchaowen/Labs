package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("10.2.8.17")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	conn.Write([]byte("hello"))

	data := make([]byte, 1024)
	n, err := conn.Read(data)

	fmt.Printf("read %s from <%s>\n", data[:n], conn.RemoteAddr())
}
