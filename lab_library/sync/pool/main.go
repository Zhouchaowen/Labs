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

// example 1
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
