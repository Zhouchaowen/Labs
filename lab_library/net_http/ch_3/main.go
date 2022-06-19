package main

import (
	"errors"
	"net/http"
)

func redirect() {
	// 重定向
	// 返回一个状态码，3xx
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 在该函数可以禁止重定向
			// return http.ErrUseLastResponse
			if len(via) > 10 {
				return errors.New("redicret to times")
			}
			return nil
		},
	}
	request, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/redirect/20",
		nil,
	)
	_, err := client.Do(request)
	if err != nil {
		panic(err)
	}
}

func main() {
	redirect()
}
