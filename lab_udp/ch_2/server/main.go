package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 8866})
	if err != nil {
		log.Fatalf("Udp Service listen report udp fail:%v", err)
	}
	defer conn.Close()
	for {
		data := make([]byte, 1024*4) // 产生大量内存消耗,引起频繁GC
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err == nil {
			// ... 做点什么
			// 如果数据处理时间过长，就会拥塞。
			// 拥塞期间若底层缓冲区满了，说不定会丢包。
			time.Sleep(10 * time.Second)
			conn.WriteToUDP(data[:n], remoteAddr)
		}
	}
}
