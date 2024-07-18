package main

import (
	"Labs/lab_go-kit/ch_5/endpoint"
	"Labs/lab_go-kit/ch_5/pb"
	"Labs/lab_go-kit/ch_5/service"
	"Labs/lab_go-kit/ch_5/transport"
	"Labs/lab_go-kit/ch_5/utils"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	utils.NewLogger("go-kit")
	golangLimit := rate.NewLimiter(1, 1) //每秒产生10个令牌,令牌桶的可以装1个令牌
	uberLimit := ratelimit.New(1)        //一秒请求一次
	server := service.NewService(utils.GetLogger())
	endpoints := endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit, uberLimit)
	grpcServer := transport.NewGRPCServer(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run :8881")
	grpcListener, err := net.Listen("tcp", ":8881")
	if err != nil {
		utils.GetLogger().Warn("Listen", zap.Error(err))
		os.Exit(0)
	}
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	pb.RegisterUserServer(baseServer, grpcServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		utils.GetLogger().Warn("Serve", zap.Error(err))
		os.Exit(0)
	}
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
