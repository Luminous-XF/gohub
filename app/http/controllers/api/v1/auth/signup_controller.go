// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/app/response"
	"gohub/pkg/jwt"

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

	userModel := user.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)

		response.CreateJSON(ctx, gin.H{
			"token": token,
			"data":  userModel,
		})
	} else {
		response.Abort500(ctx, "Failed to create user. Please try again later.")
	}
}

func (c *SignupController) SignupUsingEmail(ctx *gin.Context) {
	req := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(ctx, &req, requests.SignupUsingEmail); !ok {
		return
	}

	userModel := user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)

		response.CreateJSON(ctx, gin.H{
			"token": token,
			"data":  userModel,
		})
	} else {
		response.Abort500(ctx, "Failed to create user. Please try again later.")
	}
}
