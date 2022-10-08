package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()

	//获取当前目录下的所有文件或目录信息
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path) //打印path信息
		//fmt.Println(info.Name()) //打印文件或目录名
		return nil
	})
}
