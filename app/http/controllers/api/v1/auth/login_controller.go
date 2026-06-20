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
		return
	}

	token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
	response.JSON(ctx, gin.H{
		"token": token,
	})
}

func (c *LoginController) LoginByPassword(ctx *gin.Context) {
	req := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(ctx, &req, requests.LoginByPassword); !ok {
		return
	}

	user, err := auth.Attempt(req.LoginID, req.Password)
	if err != nil {
		response.Unauthorized(ctx, "user not found or password is incorrect")
		return
	}

	token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
	response.JSON(ctx, gin.H{
		"token": token,
	})
}

func (c *LoginController) RefreshToken(ctx *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(ctx)
	if err != nil {
		response.Error(ctx, err, "refresh token failed")
		return
	}

	response.JSON(ctx, gin.H{
		"token": token,
	})
}
