package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w871507855/BPanel/server/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			c.String(http.StatusOK, "未登录")
		}
		_, err := utils.ParseToken(auth)
		if err != nil {
			c.Abort()
			message := err.Error()
			c.JSON(http.StatusOK, message)
			return
		}
		c.Next()
	}
}
