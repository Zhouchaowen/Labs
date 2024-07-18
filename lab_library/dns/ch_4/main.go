package main

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
)

func main() {
	// 创建一个 DNS 消息对象
	m := new(dns.Msg)

	// 设置消息头部信息
	m.SetQuestion("example.com.", dns.TypeA)

	// 添加 EDNS0 首部扩展选项
	opt := new(dns.OPT)
	opt.Hdr.Name = "."
	opt.Hdr.Rrtype = dns.TypeOPT
	opt.SetUDPSize(dns.DefaultMsgSize)
	opt.SetDo()
	opt.Option = append(opt.Option, &dns.EDNS0_SUBNET{
		Code:          dns.EDNS0SUBNET,
		Family:        1, // 1 表示 IPv4，2 表示 IPv6
		SourceNetmask: 24,
		Address:       net.ParseIP("192.168.1.1").To4(), // 地址必须是 4 字节的字节数组
	})

	// 将 EDNS0 首部扩展选项添加到 DNS 消息中
	m.Extra = append(m.Extra, opt)

	// 打印 DNS 消息
	fmt.Println(m.String())
}
