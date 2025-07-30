package router

import (
	"github.com/gin-gonic/gin"
	"server/app/controllers/file"
	"server/app/controllers/system"
)

func SystemApiInit(r *gin.RouterGroup) {
	// 创建控制器实例
	userController := &system.UserController{}
	menuController := &system.MenuController{}
	MenuApiController := &system.MenuApiController{}
	roleController := &system.RoleController{}
	DepartmentController := &system.DepartmentController{}
	commonController := &system.CommonController{}
	timerController := &system.TimerController{}
	fileGroupController := &file.FileGroupController{}
	// 用户相关路由
	rUser := r.Group("/user")
	{
		//获取用户列表
		rUser.GET("/all", userController.All)
		rUser.GET("/list", userController.List)
		rUser.POST("/add", userController.Add)
		rUser.POST("/upload", userController.Up)
		rUser.DELETE("/del", userController.Dels)
		rUser.DELETE("/:id", userController.Del)
	}
	// 角色相关路由
	rRole := r.Group("/role")
	{
		//获取用户组
		rRole.GET("/group", roleController.Group)
		rRole.GET("/list", roleController.List)
		rRole.POST("/add", roleController.Add)
		rRole.PUT("/upload", roleController.Up)
		rRole.DELETE("/del", roleController.Del)
		rRole.POST("/upMenu", roleController.RoleUpMenu)
	}
	//菜单管理
	rMenu := r.Group("/menu")
	{
		//添加菜单
		rMenu.POST("/add", menuController.Add)
		//修改菜单
		rMenu.PUT("/upload", menuController.Update)
		//菜单列表
		rMenu.GET("/list", menuController.All)
		//菜单批量删除
		rMenu.DELETE("/del", menuController.Del)
		//菜单删除
		rMenu.DELETE("/:id", menuController.Del)
	}
	//部门管理
	rDepart := r.Group("/depart")
	{
		rDepart.GET("/list", DepartmentController.List)
		rDepart.POST("/add", DepartmentController.Add)
		rDepart.PUT("/upload", DepartmentController.Edit)
		rDepart.DELETE("/del", DepartmentController.Delete)
		rDepart.DELETE("/:id", DepartmentController.Del)
		rDepart.PUT("/upAdmin", DepartmentController.UserUpDepart)
		rDepart.POST("/addAdmin", DepartmentController.UserAddDepart)
	}
	//API管理
	rApi := r.Group("/api")
	{
		//添加API
		rApi.POST("/add", MenuApiController.CreateMenuApi)
		//修改API
		rApi.PUT("/upload", MenuApiController.UpdateMenuApi)
		//API列表
		rApi.GET("/list", MenuApiController.GetMenuApiList)
		//API批量删除
		rApi.DELETE("/del", MenuApiController.DeleteMenuApi)
		//API删除
		rApi.DELETE("/:id", MenuApiController.BatchDeleteMenuApi)
	}
	// 系统相关路由
	rSystem := r.Group("/system")
	{
		//文件上传
		rSystem.POST("/common/upload", commonController.UpLoad)
		//获取角色列表
		rSystem.GET("/operationLog/list", ApiControllers.OperationLogController.List)
		//获取我的菜单
		rSystem.GET("/menu/my", menuController.MenuPermissions)
	}

	// 定时任务管理
	timer := r.Group("timer")
	{
		// 任务管理器控制
		timer.POST("start", timerController.StartTimer)
		timer.POST("stop", timerController.StopTimer)
		timer.GET("status", timerController.GetTimerStatus)

		// 任务管理
		timer.GET("task/list", timerController.GetTaskList)
		timer.POST("task/create", timerController.CreateTask)
		timer.PUT("task/upload", timerController.UpdateTask)
		timer.DELETE("task/delete/:id", timerController.DeleteTask)
		timer.GET("task/get/:id", timerController.GetTask)
		timer.POST("task/execute", timerController.ExecuteTask)
		timer.POST("task/test", timerController.TestTask)
		timer.PUT("task/toggle/:id", timerController.ToggleTaskStatus)

		// 任务日志
		timer.GET("task/logs", timerController.GetTaskLogs)

		// 工具接口
		timer.GET("cron/examples", timerController.GetCronExamples)
	}
	//文件管理组
	rFileGroup := r.Group("/fileGroup")
	{
		rFileGroup.GET("/:id", fileGroupController.Index)
		rFileGroup.GET("/list", fileGroupController.List)
		rFileGroup.POST("/add", fileGroupController.Save)
		rFileGroup.PUT("/upload", fileGroupController.Edit)
		rFileGroup.DELETE("/:id", fileGroupController.Delete)
		rFileGroup.GET("/check/:id", fileGroupController.Check)
	}
}
