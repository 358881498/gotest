package main

import (
	"task4/config"
	"task4/database"
	"task4/router"
)

func main() {
	// 初始化日志
	config.InitLogger("app.log", config.INFO)
	//初始化数据库，迁移数据表
	err := database.Init()
	if err != nil {
		config.GetLogger().Fatal("数据库连接失败" + err.Error())
	}
	//初始化路由
	routers := router.Router()
	err = routers.Run(":8080")
	if err != nil {
		config.GetLogger().Fatal("路由初始化异常" + err.Error())
	}
}
