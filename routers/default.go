package routers

import "github.com/gin-gonic/gin"

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/api/v1")
	{
		user := defaultRouters.Group("/user")
		{
			user.POST("/getUser")
			user.POST("updateUser")
			user.POST("deleteUser")
			user.POST("addUser")
		}
		role := defaultRouters.Group("/role")
		{
			role.POST("/addRole")
			role.POST("/deleteRole")
			role.POST("/updateRole")
			role.POST("/getRole")
		}
		permission := defaultRouters.Group("/permission")
		{
			permission.POST("/addPermission")
			permission.POST("/deletePermission")
			permission.POST("/updatePermission")
			permission.POST("/getPermission")
		}
	}
}
