package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func LoginByPhone(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}

	msg := govalidator.MapData{
		"phone": []string{
			"required:phone number is required",
			"digits:phone number must be 11 digits",
		},
		"verify_code": []string{
			"required:verify code is required",
			"digits:verify code must be 6 digits",
		},
	}

	errs := validate(data, rules, msg)

	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type LoginByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	LoginID       string `json:"login_id" valid:"login_id"`
	Password      string `json:"password,omitempty" valid:"password"`
}

func LoginByPassword(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
	}

	msg := govalidator.MapData{
		"captcha_id": []string{
			"required:captcha id is required",
		},
		"captcha_answer": []string{
			"required:captcha answer is required",
			"digits:captcha answer must be 6 digits",
		},
		"login_id": []string{
			"required:login id is required",
			"min:login id must be at least 3 characters",
		},
		"password": []string{
			"required:password is required",
			"min:password must be at least 6 characters",
		},
	}

	errs := validate(data, rules, msg)

	_data := data.(*LoginByPasswordRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
