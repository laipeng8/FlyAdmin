// @title           CloudFile API
// @version         1.0
// @description     CloudFile 云文件管理系统 API 文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description 请输入JWT Token，格式：Bearer <token>

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"server/app"
	_ "server/docs" // 导入Swagger文档
	"server/global"
	_ "server/router"
	"syscall"
	"time"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("GetWd err:%v", err))
	}
	global.GAD_APP_PATH = dir + string(os.PathSeparator)
	app.Start()
	defer func() {
		app.DiyDefer()
	}()

	run()
}

// run 开始监听并启动web服务
func run() {
	svr := &http.Server{
		Addr:    global.Config.App.Port,
		Handler: global.GAD_R,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			global.Logger.Errorf("listen: %s\n", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
	global.Logger.Infof("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		global.Logger.Errorf("stutdown err %v", err)
	}
	global.Logger.Infof("shutdown-->ok")
}
