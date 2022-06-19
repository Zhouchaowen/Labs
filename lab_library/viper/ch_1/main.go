package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

func main() {
	v := viper.New()

	//文件的路径如何设置
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
	fmt.Printf("%s\n", v.Get("name")) // 读取指定key
	fmt.Printf("%v\n", v.Get("mysql"))
	fmt.Printf("%s\n", v.Get("mysql.host"))
	fmt.Printf("%d\n", v.Get("mysql.port"))
}
