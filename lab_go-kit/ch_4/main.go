package main

import (
	"Labs/lab_go-kit/ch_4/endpoint"
	"Labs/lab_go-kit/ch_4/service"
	"Labs/lab_go-kit/ch_4/transport"
	"Labs/lab_go-kit/ch_4/utils"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"net/http"
)

func main() {
	utils.NewLogger("go-kit")
	golangLimit := rate.NewLimiter(1, 1) //每秒产生10个令牌,令牌桶的可以装1个令牌
	uberLimit := ratelimit.New(1)        //一秒请求一次
	server := service.NewService(utils.GetLogger())
	endpoints := endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit, uberLimit)
	httpHandler := transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}

/*
	curl --location --request POST 'http://127.0.0.1:8888/login' \
	--header 'Content-Type: application/json' \
	--data-raw '{
		"account":"imianba",
		"password":"123456"
	}'

	curl --location --request GET '127.0.0.1:8888/sum?a=1&b=1' \
	--header 'Authorization: JWTToken'
*/
