package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	// 创建一个 Gin.Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":8080")
	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
}
