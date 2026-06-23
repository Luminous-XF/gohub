package middlewares

import (
	"gohub/app/response"
	"gohub/pkg/app"
	"gohub/pkg/limiter"
	"gohub/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}

	return func(ctx *gin.Context) {
		key := limiter.GetKeyIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}

	return func(ctx *gin.Context) {
		ctx.Set("limiter-once", false)

		key := limiter.GetKeyRouteWithIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

func limitHandler(ctx *gin.Context, key string, limit string) bool {
	rate, err := limiter.CheckRate(ctx, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(ctx)
		return false
	}

	ctx.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	ctx.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	ctx.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	if rate.Reached {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "Rate limit reached",
		})
		return false
	}

	return true
}
