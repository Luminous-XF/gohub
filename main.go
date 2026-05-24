package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	r := gin.New()

	// 注册中间件
	bootstrap.SetupMiddleware(r)

	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	// 运行服务
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
