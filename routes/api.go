package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello world!",
			})
		})

		authGroup := v1.Group("/auth")
		{
			ctx := new(auth.SignupController)

			authGroup.POST("/signup/phone/exist", ctx.IsPhoneExist)
			authGroup.POST("/signup/email/exist", ctx.IsEmailExist)
		}
	}
}
