package main

import (
	"Labs/lab_udp/ch_3/codec"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	_ "golang.org/x/net/bpf"
	_ "golang.org/x/net/ipv4"
)

var (
	addr = flag.String("s", "127.0.0.1", "server address")
	port = flag.Int("p", 8972, "port")
)

var (
	stat         = make(map[string]int)
	lastStatTime = int64(0)
)

func main() {
	flag.Parse()

	conn, err := net.ListenPacket("ip4:udp", *addr)
	if err != nil {
		panic(err)
	}

	cc := conn.(*net.IPConn)
	cc.SetReadBuffer(20 * 1024 * 1024)
	cc.SetWriteBuffer(20 * 1024 * 1024)

	handleConn(conn)
}

func handleConn(conn net.PacketConn) {
	fmt.Println("server ok")
	for {
		buffer := make([]byte, 1024)

		n, remoteaddr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}

		buffer = buffer[:n]

		packet := gopacket.NewPacket(buffer, layers.LayerTypeUDP, gopacket.NoCopy)

		// Get the UDP layer from this packet
		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udp, _ := udpLayer.(*layers.UDP)

			if app := packet.ApplicationLayer(); app != nil {

				data, err := codec.EncodeUDPPacket(net.ParseIP("127.0.0.1"), net.ParseIP("127.0.0.1"), uint16(udp.DstPort), uint16(udp.SrcPort), app.Payload())
				if err != nil {
					log.Printf("failed to EncodePacket: %v", err)
					return
				}

				if _, err := conn.WriteTo(data, remoteaddr); err != nil {
					log.Printf("failed to write packet: %v", err)
					conn.Close()
					return
				}
			}
		}
	}
}
