package main

import (
	"Labs/lab_grpc/ch_2/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type BidirectionalService struct {
}

func (b *BidirectionalService) BidirectionalHello(server proto.Bidirectional_BidirectionalHelloServer) error {
	defer func() {
		log.Println("客户端断开链接")
	}()
	for {
		//获取客户端信息
		recv, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println(recv)
		time.Sleep(time.Second)
		//发送服务端信息
		err = server.Send(&proto.BidirectionalResp{Message: "服务端信息"})
		if err != nil {
			return err
		}
	}
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
	proto.RegisterBidirectionalServer(grpcServer, &BidirectionalService{})
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panicf("grpcServer.Serve err: %v", err)
	}
}
