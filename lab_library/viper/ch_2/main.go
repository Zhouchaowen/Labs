package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"time"
)

//如何将线上和线下的配置文件隔离
//不用改任何代码而且线上和线上的配置文件能隔离开

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

func init() {
	os.Setenv("MXSHOP_DEBUG", "true")
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
	//刚才设置的环境变量 想要生效 我们必须得重启goland
}

func main() {
	// 通过环境变量加载配置
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("./%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("./%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
	fmt.Printf("%s\n", v.Get("name"))
	fmt.Printf("%v\n", v.Get("mysql"))
	fmt.Printf("%s\n", v.Get("mysql.host"))
	fmt.Printf("%d\n", v.Get("mysql.port"))

	//viper的功能 - 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed： ", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})

	time.Sleep(time.Second * 300)
}
