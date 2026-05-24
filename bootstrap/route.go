// Package bootstrap 用于存放项目初始化相关代码
package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoute 初始化路由
func SetupRoute(r *gin.Engine) {
	// 注册 API 路由
	routes.RegisterAPIRoutes(r)

	// 配置 404 路由
	setup404Handler(r)
}

// setup404Handler 配置 404 路由
func setup404Handler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML
			c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": http.StatusText(http.StatusNotFound),
			})
		}
	})
}
