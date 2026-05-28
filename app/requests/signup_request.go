// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupPhoneExist(data interface{}, _ *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	msg := govalidator.MapData{
		"phone": []string{
			"required:Phone number is required.",
			"digits:Phone number length must be 11 digits.",
		},
	}

	return validate(data, rules, msg)
}

func ValidateSignupEmailExist(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	msg := govalidator.MapData{
		"email": []string{
			"required: Email address is required.",
			"min:Email address length must be greater than 4 digits.",
			"max:Email address length must be less than 30 digits.",
			"email:Please enter a valid email address.",
		},
	}

	return validate(data, rules, msg)
}
