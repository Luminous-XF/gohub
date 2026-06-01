package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/app/response"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseApiController
}

func (c *VerifyCodeController) ShowCaptcha(ctx *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	logger.LogIf(err)

	response.JSON(ctx, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (c *VerifyCodeController) SendUsingPhone(ctx *gin.Context) {
	req := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(ctx, &req, requests.VerifyCodePhone); !ok {
		return
	}

	if ok := verifycode.NewVerifyCode().SendSMS(req.Phone); !ok {
		response.Abort500(ctx, "Send sms fail.")
	} else {
		response.Success(ctx)
	}
}
