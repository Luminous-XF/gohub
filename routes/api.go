package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello world!",
			})
		})

		authGroup := v1.Group("/auth")
		{
			c := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", c.IsPhoneExist)
		}
	}
}
