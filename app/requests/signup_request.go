// Package requests 处理请求数据和表单验证
package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `json:"name" valid:"name"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

type SignupUsingEmailRequest struct {
	Email           string `json:"email,omitempty" valid:"email"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `json:"name" valid:"name"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
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

func SignupUsingPhone(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	msg := govalidator.MapData{
		"phone": []string{
			"required:Phone number is required.",
			"digits:Phone number length must be 11 digits.",
		},
		"name": []string{
			"required:Name is required.",
			"alpha_num:Invalid username format. Only English letters and numbers are permitted.",
			"between:Username length must be between 3 and 20 characters.",
		},
		"password": []string{
			"required:Password is required.",
			"min:Password must be longer than 6 characters.",
		},
		"password_confirm": []string{
			"required:Password confirm is required.",
		},
		"verify_code": []string{
			"required:The verification code is required.",
			"digits:The verification code length must be 6 digits.",
		},
	}

	errs := validate(data, rules, msg)

	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

func SignupUsingEmail(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	msg := govalidator.MapData{
		"email": []string{
			"required:Email address is required.",
			"min:Email address length must be greater than 4 digits.",
			"max:Email address length must be less than 30 digits.",
			"email:Please enter a valid email address.",
			"not_exists:The email address has been registered.",
		},
		"name": []string{
			"required:Name is required.",
			"alpha_num:Invalid username format. Only English letters and numbers are permitted.",
			"between:Username length must be between 3 and 20 characters.",
		},
		"password": []string{
			"required:Password is required.",
			"min:Password must be longer than 6 characters.",
		},
		"password_confirm": []string{
			"required:Password confirm is required.",
		},
		"verify_code": []string{
			"required:The verification code is required.",
			"digits:The verification code length must be 6 digits.",
		},
	}

	errs := validate(data, rules, msg)

	_data := data.(*SignupUsingEmailRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}
