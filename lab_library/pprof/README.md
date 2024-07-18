# pprof 相关实验

- ch_1 pprof 基础用法
- ch_2 pprof 定位heap消耗
- ch_3 pprof 定位cpu消耗
- ch_4 pprof 定位gc消耗
- ch_5 pprof 定位coroutines消耗
- ch_6 pprof 定位mutex消耗
- ch_7 pprof 定位block消耗
- ch_8
- ch_9 golang 运行时指标可视化


## 代码嵌入pprof监控

```go
package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"
    "github.com/EDDYCJY/go-pprof-example/data"
)

func main() {
    go func() {
        for {
            log.Println(data.Add("https://github.com/EDDYCJY"))
        }
    }()

    http.ListenAndServe("0.0.0.0:6060", nil)
}
```

### 实时拉取远端pprof监控数据，命令行本地交互式查看

```ruby
$ go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60

Fetching profile over HTTP from http://localhost:6060/debug/pprof/profile?seconds=60
Saved profile in /Users/eddycjy/pprof/pprof.samples.cpu.007.pb.gz
Type: cpu
Duration: 1mins, Total samples = 26.55s (44.15%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 
```

> 其他选项

- cpu（CPU Profiling）: $HOST/debug/pprof/profile，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
- block（Block Profiling）：$HOST/debug/pprof/block，查看导致阻塞同步的堆栈跟踪
- goroutine：$HOST/debug/pprof/goroutine，查看当前所有运行的 goroutines 堆栈跟踪
- heap（Memory Profiling）: $HOST/debug/pprof/heap，查看活动对象的内存分配情况
- mutex（Mutex Profiling）：$HOST/debug/pprof/mutex，查看导致互斥锁的竞争持有者的堆栈跟踪
- threadcreate：$HOST/debug/pprof/threadcreate，查看创建新OS线程的堆栈跟踪

### 画图分析远端pprof监控数据

> 本地安装graphviz

```cpp
http://www.graphviz.org/download/
```

> 1.通过本地pprof文件生成图画

```go
go tool pprof -http=:8080 cpu.prof
```

> 2.通过远端pprof监控生成图画

```go
go tool pprof -http=:8080 pprof http://localhost:6060/debug/pprof/profile
```

```go
go tool pprof http://localhost:6060/debug/pprof/profile

go tool pprof http://localhost:6060/debug/pprof/heap
GODEBUG=gctrace=1 ./go-pprof-practice | grep gc

go tool pprof http://localhost:6060/debug/pprof/allocs

go tool pprof http://localhost:6060/debug/pprof/goroutine

go tool pprof http://localhost:6060/debug/pprof/mutex

go tool pprof http://localhost:6060/debug/pprof/block
```

## 参考
https://blog.wolfogre.com/posts/go-ppof-practice/

https://segmentfault.com/a/1190000019222661

https://developer.51cto.com/article/700612.html

https://learnku.com/articles/61995