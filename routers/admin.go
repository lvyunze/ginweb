package routers

import "github.com/gin-gonic/gin"

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/api/v1/admin")
	{
		adminRouters.GET("/")
	}
}
