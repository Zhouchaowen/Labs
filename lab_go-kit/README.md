# go-kit 相关实验

- ch_1 基本使用
- ch_2 添加日志中间件
- ch_3 添加JWT中间件
- ch_4 添加限流中间件
- ch_5
- ch_6
- ch_7
- ch_8
- ch_9

## 调用链
```bigquery
                             before
                               |
                               V
                        transport_decode
                               |
                               V
             —————————————————————————————————————
            |    endpoint_middleware_before       |
            |  ————————————————————————————————   |
            |  |  service_middleware_before    |  |
            |  |  ———————————————————————————  |  |
            |  |  |         service          | |  |
            |  |  ———————————————————————————  |  |
            |  |  service_middleware_after     |  |
            |   ———————————————————————————————   |                
            |     endpoint_middleware_after       |
             —————————————————————————————————————
                               |
                               V
                             after
                               |
                               V
                        transport_encode
```
before->transport_decode->endpoint_middleware_before->service_middleware_before->service->service_middleware_after->endpoint_middleware_after->transport_encode
## 参考
https://github.com/hwholiday/learning_tools/tree/master/go-kit