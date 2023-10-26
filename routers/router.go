package routers

import (
	"ginweb/handler"
	"ginweb/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	router := gin.Default()
	router.Use(middleware.CORS)
	router.Use(middleware.RequestLog)
	router.Use(handler.Recover)
	api := router.Group("/api")
	// GET：请求方式；/hello：请求的路径
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
				user.GET("/hello", func(c *gin.Context) {
					// c.JSON：返回JSON格式的数据
					if true {
						panic("异常错误")
					}
					c.JSON(200, gin.H{
						"message": "Hello world!",
					})
				})
			}
		}
		v2 := api.Group("/v2")
		{
			user := v2.Group("/user")
			{
				// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
				user.GET("/hello", func(c *gin.Context) {
					// c.JSON：返回JSON格式的数据
					c.JSON(200, gin.H{
						"message": "Hello world!",
					})
				})
			}
		}
	}

	return router
}
