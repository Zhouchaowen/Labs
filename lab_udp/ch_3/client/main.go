package main

import (
	"Labs/lab_udp/ch_3/codec"
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	conn, err := net.ListenPacket("ip4:udp", "127.0.0.1")
	if err != nil {
		panic(err)
	}

	fmt.Println("client ok")

	data, err := codec.EncodeUDPPacket(net.ParseIP("127.0.0.1"), net.ParseIP("127.0.0.1"), 8972, 0, []byte("hello"))
	if err != nil {
		log.Printf("failed to EncodePacket: %v", err)
		return
	}

	remoteAddr := &net.IPAddr{IP: net.ParseIP("127.0.0.1")}

	if _, err := conn.WriteTo(data, remoteAddr); err != nil {
		log.Printf("failed to write packet: %v", err)
		conn.Close()
		return
	}

	buffer := make([]byte, 1024)

	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		log.Fatal(err)
	}

	buffer = buffer[:n]

	packet := gopacket.NewPacket(buffer, layers.LayerTypeUDP, gopacket.NoCopy)

	// Get the UDP layer from this packet
	if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {

		if app := packet.ApplicationLayer(); app != nil {

			fmt.Printf("reply: %s\n", app.Payload())
		}
	}

}
