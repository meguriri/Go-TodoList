package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("login"); err == nil {
			c.Set("name", cookie)
			c.Next()
			return
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
	}
}
