// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{}, ctx *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required: Phone number is required.",
			"digits: Phone number length must be 11 digits.",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
