# context 相关实验

- ch_1 演示可取消上下文来防止 goroutine 泄漏。
- ch_2 演示传递了一个带有任意截止日期的上下文来告诉一个阻塞函数它应该在它到达它时立即放弃它的工作。
- ch_3 演示超时的上下文来告诉一个阻塞函数它应该在超时后放弃它的工作。
- ch_4 演示如何将一个值传递给上下文，以及如何在它存在时检索它。


## 参考

https://www.jajaldoang.com/post/golang-function-timeout-with-context/

https://www.cnblogs.com/-lee/p/12820994.html

https://codeantenna.com/a/Q47HvbRdRm

https://jasonkayzk.github.io/2021/04/21/%E4%BD%BF%E7%94%A8Uber%E5%BC%80%E6%BA%90%E7%9A%84goleak%E5%BA%93%E8%BF%9B%E8%A1%8Cgoroutine%E6%B3%84%E9%9C%B2%E6%A3%80%E6%B5%8B/

https://blog.haohtml.com/archives/19308#%E4%BA%A7%E7%94%9Fgoroutine_leak%E7%9A%84%E5%8E%9F%E5%9B%A0