package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

var HOST = "10.2.0.104"

// Watch 监听Key变动
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{HOST + ":2379"}, //如果是集群，就在后面机上所有的节点[]string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		log.Fatalf("connect to etcd failed, err:%v\n", err)
	}
	fmt.Println("connect to etcd success.")
	defer cli.Close()

	// watch key:address change 监控etcd中key的变化-创建、更改、删除
	rch := cli.Watch(context.Background(), "address") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

/*
./etcdctl --endpoints=http://10.2.0.104:2379 put address 北京市中关村
./etcdctl --endpoints=http://10.2.0.104:2379 put address 北京市中关村一街
./etcdctl --endpoints=http://10.2.0.104:2379 del address
*/
