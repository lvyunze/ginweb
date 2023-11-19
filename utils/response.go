package utils

import "github.com/gin-gonic/gin"

type MyHandler func(c *gin.Context) (string, int, interface{})

func Handler() func(h MyHandler) gin.HandlerFunc {
	return func(H MyHandler) gin.HandlerFunc {
		return func(context *gin.Context) {
			msg, code, result := H(context)
			// 调用方法
			context.JSON(200, gin.H{"msg": msg, "code": code, "result": result})

		}
	}
}
