package main

import (
	"ginweb/routers"
)

func main() {
	// 创建一个默认的路由引擎
	router := routers.InitRouter()
	router.Run()
}
