package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(middlewares.LimitIP("200-H"))

	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello world!",
			})
		})

		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))

		{
			signupController := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), signupController.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), signupController.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), signupController.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), signupController.SignupUsingEmail)

			verifyCodeController := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), verifyCodeController.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), verifyCodeController.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), verifyCodeController.SendUsingEmail)

			loginController := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), loginController.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), loginController.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), loginController.RefreshToken)

			passwordController := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), passwordController.ResetPasswordByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), passwordController.ResetPasswordByEmail)
		}
	}
}
