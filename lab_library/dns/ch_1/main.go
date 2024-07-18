package main

import (
	"bufio"
	"fmt"
	"github.com/jszwec/csvutil"
	"github.com/miekg/dns"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

// GovWebsite "序号","网站标识码","网站名称","首页网址","网站主管单位","网站状态"
type GovWebsite struct {
	Num          int    `json:"num" csv:"序号"`
	Code         string `json:"code" csv:"网站标识码"`
	Name         string `json:"name" csv:"网站名称"`
	URL          string `json:"url" csv:"首页网址"`
	Organization string `json:"organization" csv:"网站主管单位"`
	Status       string `json:"status" csv:"网站状态"`
}

var dnsServ = []string{
	"8.8.8.8",
	"1.1.1.1",
	"114.114.114.114",
	"61.139.2.69",
	"218.6.200.139",
	"182.254.116.116",
	"223.5.5.5",
	"223.6.6.6",
	"180.76.76.76",
	"101.226.4.6",
	"218.30.118.6",
	"123.125.81.6",
	"140.207.198.6",
	"223.87.238.22",
}

func main() {
	c := dns.Client{
		Timeout: 5 * time.Second,
	}

	filePath := "./gov-dm.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	dnsservLen := len(dnsServ)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var s []string
	s, _ = GetAllFile("/Users/zdns/Desktop/code_study/Labs/lab_library/dns/ch_1/csv", s)
	for _, file := range s {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("ioutil.ReadFile err: ", err)
			continue
		}
		var gws []GovWebsite
		str := strings.ReplaceAll(string(content), `"`, ``)
		err = csvutil.Unmarshal([]byte(str), &gws)
		if err != nil {
			fmt.Println("csvutil.Unmarshal err: ", err)
			continue
		}

		for _, v := range gws {
			u, err := url.Parse(strings.Trim(v.URL, " "))
			if err != nil {
				fmt.Println("url.Parse err: ", err)
				continue
			}

			m := dns.Msg{}
			// 最终都会指向一个ip 也就是typeA, 这样就可以返回所有层的cname.
			m.SetQuestion(fmt.Sprintf("%s.", u.Host), dns.TypeA)

			for try := 0; try < 3; try++ {
				dnsS := dnsServ[r.Intn(dnsservLen)]
				r, _, err := c.Exchange(&m, dnsS+":53")
				if err != nil {
					fmt.Println("c.Exchange err: ", err)
					continue
				}
				for _, ans := range r.Answer {
					fmt.Printf("%s\t%+v\n", dnsS, ans.String())
					write.WriteString(fmt.Sprintf("%+v\n", ans.String()))
				}
				//Flush将缓存的文件真正写入到文件中
				write.Flush()
				break
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}
