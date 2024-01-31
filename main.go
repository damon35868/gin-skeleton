package main

import (
	"fmt"

	"gin-skeleton/app/task"
	"gin-skeleton/config"
	router "gin-skeleton/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// config
	globalConfig, err := config.Boot()

	if err != nil {
		fmt.Println("配置失败")
	}

	// 定时任务
	task.Boot()

	router.Boot(app)
	app.Run(fmt.Sprintf(":%s", globalConfig.Server.Port))
}
