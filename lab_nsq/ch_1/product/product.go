package main

import (
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"strconv"
)

var nullLogger = log.New(ioutil.Discard, "", log.LstdFlags)

func sendMessage() {
	config := nsq.NewConfig() // 1. 创建生产者
	producer, err := nsq.NewProducer("10.2.8.17:4150", config)
	if err != nil {
		log.Fatalln("连接失败: (10.2.0.105:4150)", err)
	}

	errPing := producer.Ping() // 2. 生产者ping
	if errPing != nil {
		log.Fatalln("无法ping通: 10.2.0.105:4150", errPing)
	}

	producer.SetLogger(nullLogger, nsq.LogLevelInfo) // 3. 设置不输出info级别的日志

	for i := 0; i < 5; i++ { // 4. 生产者发布消息
		message := "消息发送测试 " + strconv.Itoa(i+10000)
		err2 := producer.Publish("one-test", []byte(message)) // 注意one-test　对应消费者consumer.go　保持一致
		if err2 != nil {
			log.Panic("生产者推送消息失败!")
		}
	}

	producer.Stop() // 5. 生产者停止执行
}

func main() {
	sendMessage()
}
