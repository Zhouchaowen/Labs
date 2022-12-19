package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"sync"
	"time"
)

// 相关配置信息
const (
	HOST = "10.2.0.104:9092"
	// Topic 注: 如果关闭了自动创建分区，使用前都需要手动创建对应分区
	Topic           = "standAlone"
	ConsumerGroupID = "cg1"
)

// sarama 库中消费者组为一个接口 sarama.ConsumerGroup 所有实现该接口的类型都能当做消费者组使用。

// MyConsumerGroupHandler 实现 sarama.ConsumerGroup 接口，作为自定义ConsumerGroup
type MyConsumerGroupHandler struct {
	name  string
	count int64
}

// Setup 执行在 获得新 session 后 的第一步, 在 ConsumeClaim() 之前
func (MyConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup 执行在 session 结束前, 当所有 ConsumeClaim goroutines 都退出时
func (MyConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim 具体的消费逻辑
func (h MyConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("[consumer] name:%s value:%s topic:%q partition:%d offset:%d\n", h.name, msg.Value, msg.Topic, msg.Partition, msg.Offset)

		// 模拟消费失败
		n, err := strconv.Atoi(string(msg.Value[14:16]))
		if err != nil || n%8 == 0 {
			return fmt.Errorf("test err")
		}

		// 标记消息已被消费 内部会更新 consumer offset
		sess.MarkMessage(msg, "")
	}
	return nil
}

func ConsumerGroup(topic, group, name string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cg, err := sarama.NewConsumerGroup([]string{HOST}, group, config)
	if err != nil {
		log.Fatal("NewConsumerGroup err: ", err)
	}
	defer cg.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler := MyConsumerGroupHandler{name: name}
		for {
			fmt.Println("running: ", name)
			/*
				![important]
				应该在一个无限循环中不停地调用 Consume()
				因为每次 Rebalance 后需要再次执行 Consume() 来恢复连接
				Consume 开始才发起 Join Group 请求 如果当前消费者加入后成为了 消费者组 leader,则还会进行 Rebalance 过程，从新分配
				组内每个消费组需要消费的 topic 和 partition，最后 Sync Group 后才开始消费
				具体信息见 https://github.com/lixd/kafka-go-example/issues/4
			*/
			err = cg.Consume(ctx, []string{topic}, handler)
			if err != nil {
				log.Println("Consume err: ", err)
			}
			<-time.After(2 * time.Second)
			// 如果 context 被 cancel 了，那么退出
			if ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()
}

// 本例展示最简单的 消费者组 的使用
/*
topic 有多个分区时，消息会自动路由到对应的分区,因为路由算法的关系 可能不会平均分.
kafka 会以分区为单位（一个分区只能被一个 consumer 消费.），让消费者组中的多个消费者消费同一个 topic 下的消息（）
如果该 topic 只有一个分区则只会有一个 consumer 能够取到消息
*/

func main() {
	ConsumerGroup(Topic, ConsumerGroupID, "C1")
}
