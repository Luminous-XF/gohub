package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/response"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"

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
