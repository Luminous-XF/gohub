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
			signupController := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", signupController.IsPhoneExist)
			authGroup.POST("/signup/email/exist", signupController.IsEmailExist)
			authGroup.POST("/signup/using-phone", signupController.SignupUsingPhone)
			authGroup.POST("/signup/using-email", signupController.SignupUsingEmail)

			verifyCodeController := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", verifyCodeController.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", verifyCodeController.SendUsingPhone)
			authGroup.POST("/verify-codes/email", verifyCodeController.SendUsingEmail)

			loginController := new(auth.LoginController)
			authGroup.POST("/login/using-phone", loginController.LoginByPhone)
		}
	}
}
