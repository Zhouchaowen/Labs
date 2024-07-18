package main

import (
	"flag"
	"net"

	"github.com/miekg/dns"
)

var (
	DstIP      string
	DomainName string
)

func main() {
	flag.StringVar(&DstIP, "dstip", "192.168.111.10:53", "Destination IP address to use in sent frames.")
	flag.StringVar(&DomainName, "domain", "asavie.com", "Domain name to use in the DNS query.")
	flag.Parse()

	query := new(dns.Msg)
	query.SetQuestion(dns.Fqdn(DomainName), dns.TypeA)
	payload, err := query.Pack()
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp", DstIP)
	if err != nil {
		panic(err)
	}

	for {
		_, err = conn.WriteTo(payload, dst)
		if err != nil {
			panic(err)
		}
	}
}
