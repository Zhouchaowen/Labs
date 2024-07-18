package main

import (
	"Labs/lab_go-kit/ch_2/endpoint"
	"Labs/lab_go-kit/ch_2/service"
	"Labs/lab_go-kit/ch_2/transport"
	"Labs/lab_go-kit/ch_2/utils"
	"net/http"
)

func main() {
	utils.NewLogger("go-kit")
	server := service.NewService(utils.GetLogger())
	endpoints := endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}

/*
	http://127.0.0.1:8888/sum?a=1&b=1
*/
