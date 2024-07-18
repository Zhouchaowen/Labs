package main

import (
	"Labs/lab_go-kit/ch_1/endpoint"
	"Labs/lab_go-kit/ch_1/service"
	"Labs/lab_go-kit/ch_1/transport"
	"fmt"
	"net/http"
)

func main() {
	server := service.NewService()
	endpoints := endpoint.NewEndPointServer(server)
	httpHandler := transport.NewHttpHandler(endpoints)
	fmt.Println("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}

/*
	http://127.0.0.1:8888/sum?a=1&b=1
*/
