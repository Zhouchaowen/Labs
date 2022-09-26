package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

const HOST = "10.2.0.104:9092"

func main() {

	config := sarama.NewConfig()

	// Return 指定将填充哪些通道。如果它们设置为 true，则必须读取它们以防止死锁。
	// 如果启用，则在使用时发生的任何错误都会在错误通道上返回（默认禁用）。
	config.Consumer.Return.Errors = true

	// Specify brokers address. This is default one
	brokers := []string{HOST}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	topic := "important"
	// How to decide partition, is it fixed value...?
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
}
