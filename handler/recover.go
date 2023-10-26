package handler

import (
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
