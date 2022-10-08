package main

import (
	"log"
	"os"
)

/**
O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
*/
func main() {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
