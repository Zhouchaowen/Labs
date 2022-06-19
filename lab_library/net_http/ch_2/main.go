package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func setGetParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	//设置参数
	params := make(url.Values)
	params.Add("name", "zcw")
	params.Add("age", "18")
	request.URL.RawQuery = params.Encode() // 等价于 http://httpbin.org/get?age=18&name=zcw

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", content)
}

/**
{
"args": {
"age": "18",
"name": "zcw"
},
"headers": {
"Accept-Encoding": "gzip",
"Host": "httpbin.org",
"User-Agent": "Go-http-client/1.1",
"X-Amzn-Trace-Id": "Root=1-601faf1a-462af8502d27a9b3490bd020"
},
"origin": "110.185.170.145",
"url": "http://httpbin.org/get?age=18&name=zcw"
}
*/

func setPostFromParams() {
	//设置参数
	// form data 形式 query string ：name=xxx&age=18
	forms := make(url.Values)
	forms.Add("name", "zcw")
	forms.Add("age", "18")
	queryforms := forms.Encode()

	r, err := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", strings.NewReader(queryforms))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

/*
{
"args": {},
"data": "",
"files": {},
"form": {
"age": "18",
"name": "zcw"
},
"headers": {
"Accept-Encoding": "gzip",
"Content-Length": "15",
"Content-Type": "application/x-www-form-urlencoded",
"Host": "httpbin.org",
"User-Agent": "Go-http-client/1.1",
"X-Amzn-Trace-Id": "Root=1-6020d6cf-5962a3753bd0282a756e2f07"
},
"json": null,
"origin": "110.185.170.145",
"url": "http://httpbin.org/post"
}
*/

func setPostJsonParams() {
	// 构建 query json
	jsons := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "zcw",
		Age:  18,
	}
	queryJsons, _ := json.Marshal(jsons)

	r, err := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(queryJsons))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

/*
{
"args": {},
"data": "{\"name\":\"zcw\",\"age\":18}",
"files": {},
"form": {},
"headers": {
"Accept-Encoding": "gzip",
"Content-Length": "23",
"Content-Type": "application/json",
"Host": "httpbin.org",
"User-Agent": "Go-http-client/1.1",
"X-Amzn-Trace-Id": "Root=1-6020d866-7e2719c44c618deb7ecdf61b"
},
"json": {
"age": 18,
"name": "zcw"
},
"origin": "110.185.170.145",
"url": "http://httpbin.org/post"
}
*/

func setHeaderParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	//设置Header参数
	request.Header.Add("user-agent", "chrome-zcw")

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", content)
}

/*
{
"args": {},
"headers": {
"Accept-Encoding": "gzip",
"Host": "httpbin.org",
"User-Agent": "chrome-zcw",
"X-Amzn-Trace-Id": "Root=1-601fb08c-3495c88b26124746001d6af7"
},
"origin": "110.185.170.145",
"url": "http://httpbin.org/get"
}
*/

func setCookies() {
	// 模拟完成登录
	// 请求页面，set-cookie，带上请求再次请求
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	request, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/cookies/set?name=zcw&age=18", nil)
	r, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	requestSecond, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/cookies", nil)
	for _, cookie := range r.Cookies() {
		requestSecond.AddCookie(cookie)
	}
	r, err = client.Do(requestSecond)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

// Response
/*
{
"cookies": {
"age": "18",
"name": "zcw"
}
}
*/

func setTimeOut() {
	// 设置代理
	//proxyUrl, _ := url.Parse("socks5://xx.xx.xx.xx:port")
	proxyUrl, _ := url.Parse("http://xx.xx.xx.xx:port")
	// 设置超时
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DialContext: func(cxt context.Context, network, addr string) (conn net.Conn, e error) {
				return net.DialTimeout(network, addr, 2*time.Second)
			},
			ResponseHeaderTimeout: 5 * time.Second,
			TLSHandshakeTimeout:   2 * time.Second,
			Proxy:                 http.ProxyURL(proxyUrl),
		},
	}
	client.Get("xxx")
}

func main() {
	setGetParams()
	setPostFromParams()
	setPostJsonParams()
	setHeaderParams()
	setCookies()
	setTimeOut()
}
