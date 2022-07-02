# concurrent 相关实验

- ch_1 通过`goroutines`并发打印字符串
- ch_2 通过`context`控制`goroutines`退出
- ch_3 `channel`基础用法
- ch_4 `channel`基础用法
- ch_5 `channel`简单生产者消费者
- ch_6 `channel`简单工作池模式
- ch_7 `channel`简单工作池模式
- ch_8 `goroutines+channel`简单服务器
- ch_9 `goroutines+channel`菊花链过滤器过程
- ch_10 `goroutines+channel`简易聊天室

# 使用
什么时候终止?
什么可以阻止它终止?

channel 非常适合用于任务分发模型 work-poll模型

buffer channel 可以减少唤醒 
close channel 必须要所有都不发送后才可以关闭，谁发送，谁关闭。
errGroup

检查竞争 go build -race main.go


## 参考
https://golangbot.com/channels/

https://www.golang-book.com/books/intro

https://github.com/yakuter/go-channels-use-cases

https://github.com/yakuter/go-concurrency

https://learnku.com/go/t/39490

https://github.com/hit9/Go-patterns-with-channel

https://writings.sh/post/goroutine-guide-part-2

https://github.com/luk4z7/go-concurrency-guide