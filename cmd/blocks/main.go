package main

import (
	"context"
	"eth/libs/config"
	"eth/libs/db"
	"eth/service"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 获取配置文件
	config.InitConfig()

	// 初始化db
	db.InitDB()

	//// 启动定时任务
	go service.TimeServ()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		zap.S().Info("timeout of 5 seconds.")
	}
	zap.S().Info("Server exiting")
}
