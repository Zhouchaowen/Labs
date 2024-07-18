package main

import (
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
	"sync"
)

var IpMap map[string]bool

func dnsQuery(wg *sync.WaitGroup, city, domain string, ip string, DnsServer string, OnlyIp bool, repeat int, v6 bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("dnsQuery:", err)
		}
		wg.Done()
	}()

	if !strings.HasSuffix(domain, ".") {
		domain += "."
	}
	c := new(dns.Client)
	m := new(dns.Msg)
	if v6 {
		m.SetQuestion(domain, dns.TypeAAAA)
	} else {
		m.SetQuestion(domain, dns.TypeA)
	}

	o := new(dns.OPT)
	o.Hdr.Name = "."
	o.Hdr.Rrtype = dns.TypeOPT
	e := new(dns.EDNS0_SUBNET) //EDNS
	e.Code = dns.EDNS0SUBNET
	if v6 {
		e.Family = 2         // 1 IPv4 2 IPv6
		e.SourceNetmask = 56 //  地址掩码 ipv4 一般为 /24  ipv6为 /56
		e.Address = net.ParseIP(ip).To16()
	} else {
		e.Family = 1
		e.SourceNetmask = 24
		e.Address = net.ParseIP(ip).To4()
	}

	e.SourceScope = 0
	o.Option = append(o.Option, e)
	m.Extra = append(m.Extra, o)
	for i := 0; i < repeat; i++ {
		in, _, err := c.Exchange(m, DnsServer) //注意:要选择支持自定义EDNS的DNS 或者是 目标NS服务器  国内DNS大部分不支持自定义EDNS数据

		if err != nil {
			log.Fatal(err)
		}
		for _, answer := range in.Answer {
			if answer.Header().Rrtype == dns.TypeA {
				fmt.Println(city, answer.(*dns.A).A.String())
			} else if answer.Header().Rrtype == dns.TypeAAAA {
				fmt.Println(city, answer.(*dns.AAAA).AAAA.String())
			}
		}
	}

}

func main() {
	Initlist()
	var domain = flag.String("d", "www.taobao.com", "domain")
	var DnsServer = flag.String("s", "8.8.8.8:53", "dns server addr")
	var ip = flag.String("ip", "", "client ip")
	var OnlyIp = flag.Bool("i", false, "Only output ip addr")
	var repeat = flag.Int("r", 1, "repeat query rounds")
	var v6 = flag.Bool("6", false, "query AAAA (ipv6)")
	flag.Parse()
	IpMap = make(map[string]bool)
	var wg sync.WaitGroup
	if (*ip != "") || (*v6) {
		*OnlyIp = true
		dnsQuery(&wg, "", *domain, *ip, *DnsServer, *OnlyIp, *repeat, *v6)
	} else {
		for city, ip := range CityMap {
			wg.Add(1)
			go dnsQuery(&wg, city, *domain, ip, *DnsServer, *OnlyIp, *repeat, *v6)
		}
	}

	wg.Wait()
}
