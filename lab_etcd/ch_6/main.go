package main

import (
	"context"
	"fmt"

	"flag"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"os"
	"os/signal"
	"time"
)

var (
	serverName = flag.String("name", "", "name of this server")
)

func main() {
	flag.Parse()
	if len(*serverName) == 0 {
		panic("server name empty")
	}

	endpoints := []string{"http://10.2.0.106:2379"}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	s1, err := concurrency.NewSession(cli)
	if err != nil {
		panic(err)
	}
	defer s1.Close()
	e1 := concurrency.NewElection(s1, "/my-election")

	go func() {
		// 开始竞选
		if err := e1.Campaign(context.Background(), *serverName); err != nil {
			fmt.Println(err)
		}

	}()

	masterName := ""

	// 每隔1s 查看一下当前的master是谁
	go func() {
		cctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		timer := time.NewTimer(time.Second)
		for range timer.C {
			timer.Reset(time.Second)
			select {
			case resp := <-e1.Observe(cctx):
				if len(resp.Kvs) > 0 {
					masterName = string(resp.Kvs[0].Value)
					fmt.Println("get master with:", masterName)
				}

			}
		}

	}()

	// 每隔5s查看一下自己是否master
	go func() {
		timer := time.NewTimer(5 * time.Second)
		for range timer.C {
			if masterName == *serverName {
				fmt.Println("oh, i'm master")
			} else {
				fmt.Println("slave!!")
			}
		}

	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	s := <-c
	fmt.Println("Got signal:", s)
	e1.Resign(context.TODO())

}
