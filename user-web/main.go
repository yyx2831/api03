package main

import (
	"api03/user-web/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	port := 8021
	//1.初始化Logger
	initialize.InitLogger()
	//2.初始化router
	Router := initialize.Routers()

	/*
		    1.S()可以获取一个全局sugar,可以让我们直接设置一个全局的logger
			2.日志级别,debug,info,warn,error,dpanic,panic,fatal
			3.S函数和L函数很用用，提供了一个全局的安全访问logger的途径
	*/

	//3.启动服务
	zap.S().Infof("启动服务器,端口号:%d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", zap.Error(err))
	}
}
