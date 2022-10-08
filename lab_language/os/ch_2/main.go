// 标准库 os.File 基础操作
package main

import (
	"fmt"
	"os"
	"time"
)

/**
// Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
func Create(name string) (file *File, err error)

// Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式
func Open(name string) (file *File, err error)

// Stat返回描述文件f的FileInfo类型值
func (f *File) Stat() (fi FileInfo, err error)

// Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo，这些FileInfo是被Lstat返回的，采用目录顺序
func (f *File) Readdir(n int) (fi []FileInfo, err error)

// Read方法从f中读取最多len(b)字节数据并写入b
func (f *File) Read(b []byte) (n int, err error)

// 向文件中写入字符串
func (f *File) WriteString(s string) (ret int, err error)

// Sync递交文件的当前内容进行稳定的存储。一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存
func (f *File) Sync() (err error)

// Close关闭文件f，使文件不能用于读写
func (f *File) Close() error
*/
func main() {
	// 获取当前目录
	dir, err := os.Getwd()
	fmt.Println(dir, err)

	file := dir + "/tmp.txt"
	var fh *os.File

	// 读取文件状态
	fi, _ := os.Stat(file)
	if fi == nil {
		fh, _ = os.Create(file) // 文件不存在就创建
	} else {
		fh, _ = os.OpenFile(file, os.O_RDWR, 0666) // 文件存在就打开
	}

	w := []byte("hello go language" + time.Now().String())
	n, err := fh.Write(w)
	fmt.Println(n, err)

	// 设置下次读写位置
	ret, err := fh.Seek(0, 0)
	fmt.Println("当前文件指针位置", ret, err)

	b := make([]byte, 128)
	n, err = fh.Read(b)
	fmt.Println(n, err, string(b))

	fh.Close()
}
