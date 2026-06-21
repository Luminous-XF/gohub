package middlewares

import (
	"fmt"
	"gohub/app/models/user"
	"gohub/app/response"
	"gohub/pkg/config"
	"gohub/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := jwt.NewJWT().ParseToken(ctx)
		if err != nil {
			response.Unauthorized(ctx, fmt.Sprintf("Please refer to the API authentication documentation for %v.", config.GetString("app.name")))
			return
		}

		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(ctx, "No matching user exists; this user account may have been removed.")
			return
		}

		ctx.Set("current_user_id", userModel.GetStringID())
		ctx.Set("current_user_name", userModel.Name)
		ctx.Set("current_user", userModel)

		ctx.Next()
	}
}
