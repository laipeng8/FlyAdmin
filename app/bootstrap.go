package app

import (
	"os"
	"server/global"
	"server/initialize"
	"server/pkg/timer"
	"server/router"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Start() {
	//关闭GIN-debug
	gin.SetMode(gin.ReleaseMode)
	global.SuperAdmin = "administrator"
	global.GAD_R = gin.Default()
	global.Config = initialize.ConfigInit(global.GAD_APP_PATH)
	loadObject()
	router.RouteInit(global.GAD_R)

	// 启动定时任务管理器
	//startTimerManager()
}

func TestLoad() {
	dir, err := os.Getwd()
	if err != nil {
	}
	global.GAD_APP_PATH = dir + "/../"
	global.Config = initialize.ConfigInit(global.GAD_APP_PATH)
	loadObject()
	startTimerManager()
}

func loadObject() {
	global.Db = initialize.DbInit(global.Config)
	//initialize.AutoMigrate(global.Db) //自动迁移数据库
	global.Logger = initialize.ZapInit(global.Config)
	global.EventDispatcher = initialize.EventInit()
	global.Limiter = rate.NewLimiter(global.Config.Rate.Limit, global.Config.Rate.Burst)
	global.ValidatorManager = initialize.InitValidator()

}

func startTimerManager() {
	// 启动定时任务管理器
	taskManager := timer.GetTaskManager()
	if err := taskManager.Start(); err != nil {
		global.Logger.Errorf("启动定时任务管理器失败: %v", err)
	} else {
		global.Logger.Info("定时任务管理器启动成功")
	}
}

func DiyDefer() {
	// 停止定时任务管理器
	taskManager := timer.GetTaskManager()
	if err := taskManager.Stop(); err != nil {
		global.Logger.Errorf("停止定时任务管理器失败: %v", err)
	}

	initialize.DbClose(global.Db)
	initialize.ZapSync(global.Logger)
}
