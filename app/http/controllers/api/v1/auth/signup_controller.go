// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"fmt"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

func (c *SignupController) IsPhoneExist(ctx *gin.Context) {
	// 初始化请求对象
	req := requests.SignupPhoneExistRequest{}

	// 解析 JSON 请求
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 解析失败, 返回 422 状态码和错误信息
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		// 打印错误信息
		fmt.Println(err.Error())

		// 出错了, 中断请求
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&req, ctx)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(req.Phone),
	})
}
