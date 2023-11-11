package routers

import (
	"ginweb/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"msg":     err,
				"success": false,
			})
			c.Abort()
		}
	}()
	c.Next()
}

func InitRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	router := gin.Default()
	router.Use(middleware.CORS)
	router.Use(middleware.RequestLog)
	router.Use(Recover)
	AdminRoutersInit(router)
	DefaultRoutersInit(router)
	return router
}
