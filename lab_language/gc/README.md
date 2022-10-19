# GC 相关实验
> GC 扫描会忽略基础类型组成的一维数组,map
> 
> 基础类型：int,bool,float,double
> 
> 引用类型：struct,slice,string
> 
> 注意： 基础类型的array总内存小于128字节的情况下是值类型 ([128]byte,[16]int,[64]int16,[32]int32,[16]int64,[128]bool,[32]float32,[16]float64)

- ch_1 基础类型Slice Gc问题  
- ch_2 结构体类型 Gc问题
- ch_3 go_issuse_9477 实现
- ch_4 Map类型 Gc 问题
- ch_5 Mmap分配内存绕过Gc
- ch_6 Map中值类型的Key/Value超过128字节后会变成指针

## 参考
https://mp.weixin.qq.com/s/jGGCccMOx4s5asG2IXWNMQ
