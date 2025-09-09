package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	// 配置初始化参数, 依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件, 如 --env=testing 则加载 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 创建一个 Gin.Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.GetWithDefault[string]("app.port", "8080"))
	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
}
