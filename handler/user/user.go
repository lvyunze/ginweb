package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUsersHandler(c *gin.Context) {
	// 在这里编写处理程序的逻辑
	users := []string{"User1", "User2", "User3"}

	// 假设我们从数据库或其他数据源中获取了用户列表
	// 这里只是一个示例，实际应用中可能需要进行更复杂的数据操作

	// 返回 JSON 响应
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
