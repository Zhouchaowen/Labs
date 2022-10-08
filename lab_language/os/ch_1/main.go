// 标准库 os 基础操作
package main

import (
	"fmt"
	"os"
)

/**
func Hostname() (name string, err error) 		// Hostname返回内核提供的主机名
func Environ() []string 						// Environ返回表示环境变量的格式为"key=value"的字符串的切片拷贝
func Getenv(key string) string 					// Getenv检索并返回名为key的环境变量的值
func Getpid() int 								// Getpid返回调用者所在进程的进程ID
func Exit(code int) 							// Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行
func Stat(name string) (fi FileInfo, err error) // 获取文件信息
func Getwd() (dir string, err error) 			// Getwd返回一个对应当前工作目录的根路径
func Mkdir(name string, perm FileMode) error 	// 使用指定的权限和名称创建一个目录
func MkdirAll(path string, perm FileMode) error // 使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误
func Remove(name string) error 					// 删除name指定的文件或目录
func TempDir() string 							// 返回一个用于保管临时文件的默认目录
var Args []string Args保管了命令行参数，第一个是程序名。
*/
func main() {
	// 预定义变量, 保存命令行参数
	fmt.Println("Args: ", os.Args)

	// 获取host name
	hostname, err := os.Hostname()
	fmt.Println("Hostname: ", hostname)
	fmt.Println("Getpid: ", os.Getpid())

	// 获取全部环境变量
	env := os.Environ()
	for k, v := range env {
		fmt.Println("Environ: ", k, v)
	}

	// 设置/获取指定环境变量
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	os.Setenv("SOME_KEY", "value")
	os.Setenv("EMPTY_KEY", "")

	show("SOME_KEY")
	show("EMPTY_KEY")
	show("MISSING_KEY")

	// 终止程序
	// os.Exit(1)

	// 获取一条环境变量
	fmt.Println("Getenv: ", os.Getenv("PATH"))

	// 获取当前目录
	dir, err := os.Getwd()
	fmt.Println("Getwd: ", dir, err)

	// 创建目录
	err = os.Mkdir(dir+"/new_file", 0755)
	fmt.Println(err)

	// 创建目录
	err = os.MkdirAll(dir+"/new", 0755)
	fmt.Println(err)

	// 删除目录
	err = os.Remove(dir + "/new_file")
	err = os.Remove(dir + "/new")
	fmt.Println(err)

	// 创建临时目录
	tmp_dir := os.TempDir()
	fmt.Println(tmp_dir)
}
