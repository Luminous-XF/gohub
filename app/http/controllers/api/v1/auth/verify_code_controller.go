package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseApiController
}

func (c *VerifyCodeController) ShowCaptcha(ctx *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	logger.LogIf(err)

	ctx.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
