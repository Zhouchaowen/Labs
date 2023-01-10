package main

import (
	"context"
	"flag"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"time"
)

const prefix = "/election-demo"
const prop = "local"

var leaderFlag bool
var NodeName string

func init() {
	flag.StringVar(&NodeName, "n", "master", "node name")
}

func main() {
	flag.Parse()

	endpoints := []string{"10.2.0.106:2379"}
	donec := make(chan struct{})

	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	go campaign(cli, prefix, prop)

	go func() {
		ticker := time.NewTicker(time.Duration(5) * time.Second)
		for {
			select {
			case <-ticker.C:
				doCrontab()
			}
		}
	}()

	<-donec
}

func campaign(c *clientv3.Client, election string, prop string) {
	for {
		s, err := concurrency.NewSession(c, concurrency.WithTTL(15))
		if err != nil {
			fmt.Println(err)
			continue
		}
		e := concurrency.NewElection(s, election)
		ctx := context.TODO()

		fmt.Println("ready lock")
		if err = e.Campaign(ctx, prop); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("elect: success")
		leaderFlag = true

		select {
		case <-s.Done():
			leaderFlag = false
			fmt.Println("elect: expired")
		}
	}
}

func doCrontab() {
	if leaderFlag == true {
		fmt.Println(NodeName, "doCrontab")
	}
}
