// pool 的问题
// 1.内存泄漏-比如 encoding、json 中类似的问题：将容量已经变得很大的 Buffer 再放回 Pool 中，
// 导致内存泄漏。后来在元素放回时，增加了检查逻辑，改成放回的超过一定大小的 buffer，就直接丢弃掉，不再放到池子中。
//
// 2.内存浪费-要做到物尽其用，尽可能不浪费的话，我们可以将 buffer 池分成几层。
// 首先，小于 512 byte 的元素的 buffer 占一个池子；其次，小于 1K byte 大小的元素占一个池子；
// 再次，小于 4K byte 大小的元素占一个池子。这样分成几个池子以后，就可以根据需要，到所需大小的池子中获取 buffer 了。
//（net/http/server.go）
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		// Pool 的 New 函数通常应该只返回指针类型，因为可以将指针放入返回接口值中而无需分配
		return new(bytes.Buffer)
	},
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(time.Now().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main() {
	Log(os.Stdout, "path", "/search?q=flowers")
	// Output: 2006-01-02T15:04:05Z path=/search?q=flowers

	// example 2
	fmt.Println()
	s := sync.Pool{}
	s.Put("test")
	fmt.Println(s.Get())
}
