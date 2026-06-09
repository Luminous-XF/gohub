// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/app/response"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

func (c *SignupController) IsPhoneExist(ctx *gin.Context) {
	req := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(ctx, &req, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	response.JSON(ctx, gin.H{
		"exist": user.IsPhoneExist(req.Phone),
	})
}

func (c *SignupController) IsEmailExist(ctx *gin.Context) {
	req := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(ctx, &req, requests.ValidateSignupEmailExist); !ok {
		return
	}

	response.JSON(ctx, gin.H{
		"exist": user.IsEmailExist(req.Email),
	})
}

func (c *SignupController) SignupUsingPhone(ctx *gin.Context) {
	req := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(ctx, &req, requests.SignupUsingPhone); !ok {
		return
	}

	_user := user.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		response.CreateJSON(ctx, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(ctx, "Failed to create user. Please try again later.")
	}
}
