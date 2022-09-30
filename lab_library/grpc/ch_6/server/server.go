package main

import (
	"Labs/lab_library/grpc/ch_6/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	//time.Sleep(2 * time.Second)
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 拦截函数
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求") // 拦截前
		res, err := handler(ctx, req)
		fmt.Println("请求已经完成") // 拦截后
		return res, err
	}

	// 注册拦截器
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)

	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
