package main

import (
	"Labs/lab_grpc/ch_8/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ServerSideService struct {
}

func (s *ServerSideService) ServerSideHello(request *proto.ServerSideRequest, server proto.ServerSide_ServerSideHelloServer) error {
	log.Println(request.Name)
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := server.Send(&proto.ServerSideResp{Message: "你好"})
		if err != nil {
			return err
		}
	}
	return nil
}

const (
	// Address 监听地址
	Address string = ":8546"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Panic("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	proto.RegisterServerSideServer(grpcServer, &ServerSideService{})
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("grpcServer.Serve err: %v", err)
	}
}
