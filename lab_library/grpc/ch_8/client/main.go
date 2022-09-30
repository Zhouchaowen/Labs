package main

import (
	"Labs/lab_library/grpc/ch_8/proto"
	"context"
	"google.golang.org/grpc"
	"io"
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
	grpcClient := proto.NewServerSideClient(conn)
	// 创建发送结构体
	req := proto.ServerSideRequest{
		Name: "我来打开你啦",
	}
	//获取流
	stream, err := grpcClient.ServerSideHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call SayHello err: %v", err)
	}
	for n := 0; n < 5; n++ {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Message)
	}
}
