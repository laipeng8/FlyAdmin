package middleware

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/global/response"
)

func EnvCheck() gin.HandlerFunc {

	return func(context *gin.Context) {

		if global.Config.Env == "dev" {
			context.Next()
		} else {
			response.Failed(context, "生产环境当前操作不允许！")
			context.Abort()
			return
		}
	}
}
