package main

import (
	"Labs/lab_library/grpc/ch_3/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

const (
	// ServerAddress 连接地址
	ServerAddress string = ":8546"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := proto.NewClientSideClient(conn)
	// 创建发送结构体
	res, err := grpcClient.ClientSideHello(context.Background())
	if err != nil {
		log.Fatalf("Call SayHello err: %v", err)
	}
	for i := 0; i < 5; i++ {
		//通过 Send方法发送流信息
		err = res.Send(&proto.ClientSideRequest{Name: "客户端流式"})
		if err != nil {
			return
		}
	}
	// 打印返回值
	log.Println(res.CloseAndRecv())
}
