package main

import (
	"flag"
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

// Config contains program configuration options.
var Config struct {
	Url      string
	Filename string
}

// init is called before main.
func init() {
	flag.StringVar(&Config.Url, "u", "", "request url")
	flag.StringVar(&Config.Filename, "f", "tmp.jpg", "save file name")
	flag.Parse()
}

func main() {
	// 自动下载文件
	downloadFile(Config.Url, Config.Filename)
}
