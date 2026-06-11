package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/app/response"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseApiController
}

func (c *LoginController) LoginByPhone(ctx *gin.Context) {
	req := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(ctx, &req, requests.LoginByPhone); !ok {
		return
	}

	user, err := auth.LoginByPhone(req.Phone)
	if err != nil {
		response.Error(ctx, err, "user not found")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(ctx, gin.H{
			"token": token,
		})
	}
}
