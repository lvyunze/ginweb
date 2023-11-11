package routers

import "github.com/gin-gonic/gin"

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/api/v1/default")
	{
		defaultRouters.GET("/")
	}
}
