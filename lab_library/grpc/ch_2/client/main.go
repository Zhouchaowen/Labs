package main

import (
	"Labs/lab_grpc/ch_2/proto"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
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
	grpcClient := proto.NewBidirectionalClient(conn)
	//获取流信息
	stream, err := grpcClient.BidirectionalHello(context.Background())
	if err != nil {
		log.Fatalf("get BidirectionalHello stream err: %v", err)
	}

	for n := 0; n < 5; n++ {
		err := stream.Send(&proto.BidirectionalRequest{Name: "双向流 rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
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
