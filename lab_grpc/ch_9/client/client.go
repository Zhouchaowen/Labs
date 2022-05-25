package main

import (
	"Labs/lab_grpc/ch_9/pkg/auth"
	"Labs/lab_grpc/ch_9/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// Address 连接地址
const Address string = ":8000"

var grpcClient proto.SimpleClient

// route 调用服务端Route方法
func route() {
	// 创建发送结构体
	req := proto.SimpleRequest{
		Data: "grpc",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}

func main() {
	//从输入的证书文件中为客户端构造TLS凭证。  CA 证书来校验服务端的证书有效性，生成服务端证书时指定的CN参数
	creds, err := credentials.NewClientTLSFromFile("../pkg/tls/ca.crt", "www.lixueduan.com")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	//构建Token
	token := auth.Token{
		AppID:     "grpc_token",
		AppSecret: "123456",
	}
	// 连接服务器 设置tls 设置认证token
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = proto.NewSimpleClient(conn)
	route()
}
