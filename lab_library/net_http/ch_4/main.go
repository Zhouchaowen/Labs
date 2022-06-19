package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	fmt.Printf("\r当前进度：%.2f%%", float64(r.Current*10000/r.Total)/100)
	return
}

func downloadFile(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 实现进度显示
	reader := &Reader{
		Reader: r.Body,
		Total:  r.ContentLength,
	}

	n, err := io.Copy(f, reader)
	fmt.Println(n, err)
}

func main() {
	// 自动下载文件
	downloadFile("https://user-gold-cdn.xitu.io/2019/6/30/16ba8cb6465a6418?w=826&h=782&f=png&s=279620", "filename.jpg")
}
