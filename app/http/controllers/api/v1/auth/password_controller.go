package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/app/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseApiController
}

func (c *PasswordController) ResetPasswordByPhone(ctx *gin.Context) {
	req := requests.ResetPasswordByPhoneRequest{}
	if ok := requests.Validate(ctx, &req, requests.ResetPasswordByPhone); !ok {
		return
	}

	userModel := user.GetByPhone(req.Phone)
	if userModel.ID == 0 {
		response.Abort404(ctx)
	}

	userModel.Password = req.Password
	userModel.Save()
	response.Success(ctx)
}

func (c *PasswordController) ResetPasswordByEmail(ctx *gin.Context) {
	req := requests.ResetPasswordByEmailRequest{}
	if ok := requests.Validate(ctx, &req, requests.ResetPasswordByEmail); !ok {
		return
	}

	userModel := user.GetByEmail(req.Email)
	if userModel.ID == 0 {
		response.Abort404(ctx)
		return
	}

	userModel.Password = req.Password
	userModel.Save()
	response.Success(ctx)
}
