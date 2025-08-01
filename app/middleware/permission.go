package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/app/models"
	"server/app/repositorys"
	"server/global"
)

func Permission() gin.HandlerFunc {

	return func(c *gin.Context) {

		var (
			apiList              []models.MenuApiList
			systemMenuRepository repositorys.SystemMenuRepository
		)

		err := systemMenuRepository.GetApiList(c, &apiList)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "无权限访问!",
			})
			c.Abort()
			return
		}

		isAllow := false
		for _, api := range apiList {
			if api.Url == c.Request.URL.Path {
				isAllow = true
			}
		}
		Claims, _ := systemMenuRepository.GetCustomClaims(c)

		if !isAllow && !global.IsSuperAdmin(Claims.Roles, global.SuperAdmin) {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "无权限访问!",
			})
			c.Abort()
			return
		}
		c.Next()

	}
}
