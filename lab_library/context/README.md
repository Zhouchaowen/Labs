# context 相关实验

- ch_1 `WithCancel`演示可取消上下文来防止 goroutine 泄漏。
- ch_2 `WithDeadline`演示传递了一个带有任意截止日期的上下文来告诉一个阻塞函数它应该在它到达它时立即放弃它的工作。
- ch_3 `WithTimeout`演示超时的上下文来告诉一个阻塞函数它应该在超时后放弃它的工作。
- ch_4 `WithValue`演示如何将一个值传递给上下文，以及如何在它存在时检索它。



## Notes
- Incoming requests to a server should create a Context.
- Outgoing calls to servers should accept a Context.
- The chain of function calls between them must propagate the Context.
- Replace a Context using WithCancel, WithDeadline, WithTimeout, or WithValue.
- When a Context is canceled, all Contexts derived from it are also canceled.
- Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
- Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.
- Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
- The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.

context作用域是请求级别，通过链式结构，将每个不同域区别开

超时处理和取消，存取元数据(只能做旁路数据挂载)

往下传的时候如果要修改挂载的元数据，必须 copy on write

## 参考

https://www.jajaldoang.com/post/golang-function-timeout-with-context/

https://www.cnblogs.com/-lee/p/12820994.html

https://codeantenna.com/a/Q47HvbRdRm

https://jasonkayzk.github.io/2021/04/21/%E4%BD%BF%E7%94%A8Uber%E5%BC%80%E6%BA%90%E7%9A%84goleak%E5%BA%93%E8%BF%9B%E8%A1%8Cgoroutine%E6%B3%84%E9%9C%B2%E6%A3%80%E6%B5%8B/

https://blog.haohtml.com/archives/19308#%E4%BA%A7%E7%94%9Fgoroutine_leak%E7%9A%84%E5%8E%9F%E5%9B%A0