package main

import (
	"log"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

// 相关配置信息
const (
	HOST = "10.2.0.104:9092"
	// Topic 注: 如果关闭了自动创建分区，使用前都需要手动创建对应分区
	Topic            = "standAlone"
	Topic2           = "consumerGroup"
	Topic3           = "benchmark"
	TopicPartition   = "partition"
	TopicCompression = "compression"
	DefaultPartition = 0
	ConsumerGroupID  = "cg1"
	ConsumerGroupID2 = "cg2"
)

/*
	本例展示最简单的 同步生产者 的使用（除同步生产者外 kafka 还有异步生产者）
	名词 sync producer
*/

func Producer(topic string, limit int) {
	config := sarama.NewConfig()
	// 同步生产者必须同时开启 Return.Successes 和 Return.Errors
	// 因为同步生产者在发送之后就必须返回状态，所以需要两个都返回
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true // 这个默认值就是 true 可以不用手动 赋值

	producer, err := sarama.NewSyncProducer([]string{HOST}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()

	var successes, errors int
	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg) // 发送逻辑也是封装的异步发送逻辑，可以理解为将异步封装成了同步
		if err != nil {
			log.Printf("SendMessage:%d err:%v\n ", i, err)
			errors++
			continue
		}
		successes++
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, str)
	}
	log.Printf("发送完毕 总发送条数:%d successes: %d errors: %d\n", limit, successes, errors)
}

func main() {
	Producer(Topic, 100)
}
