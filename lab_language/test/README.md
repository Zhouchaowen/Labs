# test 相关用法实验

- ch_1 单元测试和基准测试
- ch_2 并行测试
- ch_3 竞争检查
- ch_4 代码覆盖率

## 基础命令
```bash
zcw% go test -H
  -test.bench regexp [-bench=xxxx]
        只运行与regexp匹配的基准测试 run only benchmarks matching regexp
  -test.benchmem     [-benchmem]
        打印基准测试的内存分配 print memory allocations for benchmarks
  -test.benchtime d  [-benchtime 10]
        运行每个基准测试，持续时间为d(默认为1秒) run each benchmark for duration d (default 1s)
  -test.blockprofile file
        将goroutine阻塞配置文件写入文件 write a goroutine blocking profile to file
  -test.blockprofilerate rate
        设置阻塞 profile 文件速率(参见runtime.SetBlockProfileRate)(默认为1) set blocking profile rate (see runtime.SetBlockProfileRate) (default 1)
  -test.count n
        运行测试和基准测试n次(默认1次) run tests and benchmarks n times (default 1)
  -test.coverprofile file
        写一个覆盖 profile 文件 write a coverage profile to file
  -test.cpu list
        用逗号分隔的运行每个测试的CPU计数列表 comma-separated list of cpu counts to run each test with
  -test.cpuprofile file
        将CPU profile 文件写入文件 write a cpu profile to file
  -test.failfast
        在第一次测试失败后不开始新的测试 do not start new tests after the first test failure
  -test.list regexp
        列出符合regexp的测试、示例和基准测试，然后退出 list tests, examples, and benchmarks matching regexp then exit
  -test.memprofile file
        将分配配置文件写入文件 write an allocation profile to file
  -test.memprofilerate rate
        设置内存分配分析速率(参见runtime.MemProfileRate) set memory allocation profiling rate (see runtime.MemProfileRate)
  -test.mutexprofile string
        执行后将互斥锁争用配置文件写入命名文件 write a mutex contention profile to the named file after execution
  -test.mutexprofilefraction int
        如果>= 0，调用runtime.SetMutexProfileFraction()(默认为1) if >= 0, calls runtime.SetMutexProfileFraction() (default 1)
  -test.outputdir dir
        将配置文件写入目录 write profiles to dir
  -test.paniconexit0
        panic on call to os.Exit(0)
  -test.parallel n
        并行运行最多n个测试(默认为4个) run at most n tests in parallel (default 4)
  -test.run regexp
        只运行与regexp匹配的测试和示例 run only tests and examples matching regexp
  -test.short
        运行更小的测试套件以节省时间 run smaller test suite to save time
  -test.shuffle string
        randomize the execution order of tests and benchmarks (default "off")
  -test.testlogfile file
        write test action log to file (for use only by cmd/go)
  -test.timeout d
        持续时间d(默认为0，禁用超时)后的恐慌测试二进制文件 panic test binary after duration d (default 0, timeout disabled)
  -test.trace file
        将执行跟踪写入文件 write an execution trace to file
  -test.v
        verbose: print additional output
```

## 参考
https://www.jianshu.com/p/595eabe003c9

https://brantou.github.io/2017/05/24/go-cover-story/

https://yizhi.ren/2019/06/15/gocoverage/

https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package

https://quii.gitbook.io/learn-go-with-tests/

https://blog.logrocket.com/a-deep-dive-into-unit-testing-in-go/