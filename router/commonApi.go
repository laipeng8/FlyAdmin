package router

import (
	"server/app/controllers/file"

	"github.com/gin-gonic/gin"
)

func CommonApiInit(r *gin.RouterGroup) {
	var commonController = ApiControllers.CommonController
	// 文件管理
	fileController := file.FileController{}
	{

		//登录
		r.POST("/user/login", ApiControllers.UserController.Login)
		//静态文件
		r.Static("/system/common/file/upload", commonController.GetFileBasePath())

	}
	//文件管理器接口
	{
		r.POST("/file/upload", fileController.Upload)
		r.PUT("/file/edit", fileController.Edit)
		r.GET("/file/:id", fileController.Index)
		r.DELETE("/file/:id", fileController.Delete)
		r.DELETE("/file/batch", fileController.BatchDelete)
	}
}
