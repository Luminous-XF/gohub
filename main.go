package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusText(http.StatusNotFound),
				"message": http.StatusText(http.StatusNotFound),
			})
		}
	})

	// 运行服务
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
