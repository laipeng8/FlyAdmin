package router

import (
	"server/app/controllers/file"
	"server/app/controllers/system"
	"server/app/middleware"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	system.UserController
	system.CommonController
	system.MenuController
	system.RoleController
	system.OperationLogController
	file.FileController
}

var ApiControllers = new(Controllers)

func RouteInit(e *gin.Engine) {
	// 全局中间件
	e.Use(middleware.Cors())

	// 不需要 JWT 验证的路由组
	w := e.Group("api")
	w.Use(middleware.Cors(), middleware.Limiter(), middleware.Event())
	// 初始化不需要 JWT 验证的接口
	CommonApiInit(w)

	// 需要 JWT 验证的路由组
	r := e.Group("api")
	//r.Use(middleware.JWTAuth(), middleware.Permission(), middleware.OperationLog())
	r.Use(middleware.JWTAuth())

	// 初始化需要 JWT 验证的接口
	SystemApiInit(r)
}
