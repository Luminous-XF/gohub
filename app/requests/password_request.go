package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetPasswordByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetPasswordByPhone(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}

	msg := govalidator.MapData{
		"phone": []string{
			"required:phone number is required",
			"digits:phone number must be 6 digits",
		},
		"verify_code": []string{
			"required:verify code is required",
			"digits:verify code must be 6 digits",
		},
		"password": []string{
			"required:password is required",
			"min:password must be at least 6 characters",
		},
	}

	errs := validate(data, rules, msg)

	_data := data.(*ResetPasswordByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}
