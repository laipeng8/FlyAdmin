package middleware

import (
	"github.com/gin-gonic/gin"
	"server/global"
)

func Event() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("e", global.EventDispatcher)
		c.Next()
	}
}
