package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() {
	r, err := http.Get("http://httpbin.org/get")
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

// response
/*
{
"args": {},
"headers": {
"Accept-Encoding": "gzip",
"Host": "httpbin.org",
"User-Agent": "Go-http-client/1.1",
"X-Amzn-Trace-Id": "Root=1-601fa318-5d6fb6a826a2e7c850cf4fbb"
},
"origin": "110.185.170.145",
"url": "http://httpbin.org/get"
}
*/

func post() {
	r, err := http.Post("http://httpbin.org/post", "", nil)
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

// response
/*
{
"args": {},
"data": "",
"files": {},
"form": {},
"headers": {
"Accept-Encoding": "gzip",
"Content-Length": "0",
"Host": "httpbin.org",
"User-Agent": "Go-http-client/1.1",
"X-Amzn-Trace-Id": "Root=1-601fa343-63ce306a26781d7170d63922"
},
"json": null,
"origin": "110.185.170.145",
"url": "http://httpbin.org/post"
}
*/

// net/http没有提供现成的Put方法，只能自己封装
func put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
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

// response
/*
{
"args": {},
"data": "",
"files": {},
"form": {},
"headers": {
"Accept-Encoding": "gzip",
"Content-Length": "0",
"Host": "httpbin.org",
"User-Agent": "Go-http-client/1.1",
"X-Amzn-Trace-Id": "Root=1-601fa4b2-67e9281e60b5273150213409"
},
"json": null,
"origin": "110.185.170.145",
"url": "http://httpbin.org/put"
}
*/

// net/http没有提供现成的Delete方法，只能自己封装
func delete() {
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
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

// response
/*
{
"args": {},
"data": "",
"files": {},
"form": {},
"headers": {
"Accept-Encoding": "gzip",
"Host": "httpbin.org",
"User-Agent": "Go-http-client/1.1",
"X-Amzn-Trace-Id": "Root=1-601fa58d-544b327f7050ea42060d5e83"
},
"json": null,
"origin": "110.185.170.145",
"url": "http://httpbin.org/delete"
}
*/

func main() {
	get()
	post()
	put()
	delete()
}
