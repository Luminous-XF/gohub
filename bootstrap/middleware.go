package bootstrap

import (
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupMiddleware 配置中间件
func SetupMiddleware(r *gin.Engine) {
	registerGlobalMiddleware(r)
}

// registerGlobalMiddleware 配置全局中间件
func registerGlobalMiddleware(r *gin.Engine) {
	r.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}
