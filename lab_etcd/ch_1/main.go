package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

var HOST = "10.2.0.104"

// PUT/GET 基础操作
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{HOST + ":2379"}, //如果是集群，就在后面机上所有的节点[]string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("connect to etcd failed, err:%v\n", err)
	}
	fmt.Println("connect to etcd success.")
	defer cli.Close()

	// put值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "address", "北京市")
	cancel()
	if err != nil {
		log.Fatalf("put to etcd failed, err:%v\n", err)
	}

	//context超时控制
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "address")
	cancel()
	if err != nil {
		log.Fatalf("get from etcd failed,err %v\n", err)
	}

	//遍历键值对
	for _, kv := range resp.Kvs {
		fmt.Printf("%s:%s \n", kv.Key, kv.Value)
	}
}
