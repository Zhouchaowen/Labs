package main

import (
	"Labs/lab_library/grpc/ch_3/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ClientSideService struct {
}

func (c *ClientSideService) ClientSideHello(server proto.ClientSide_ClientSideHelloServer) error {
	for i := 0; i < 5; i++ {
		recv, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println("客户端信息：", recv)
	}
	//服务端最后一条消息发送
	err := server.SendAndClose(&proto.ClientSideResp{Message: "关闭"})
	if err != nil {
		return err
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
		log.Panicf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	proto.RegisterClientSideServer(grpcServer, &ClientSideService{})
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panicf("grpcServer.Serve err: %v", err)
	}
}
