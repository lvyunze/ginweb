package main

import (
	"ginweb/routers"
	"github.com/spf13/viper"
)

func main() {
	// 创建一个默认的路由引擎
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	// 注册每次配置文件发生变更后都会调用的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed: %s\n", e.Name)
	})
	// 监控并重新读取配置文件，需要确保在调用前添加了所有的配置路径
	viper.WatchConfig()
	router := routers.InitRouter()
	router.Run()
}
