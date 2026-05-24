package bootstrap

import "github.com/gin-gonic/gin"

// SetupMiddleware 配置中间件
func SetupMiddleware(r *gin.Engine) {
	registerGlobalMiddleware(r)
}

// registerGlobalMiddleware 配置全局中间件
func registerGlobalMiddleware(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}
