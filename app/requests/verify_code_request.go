package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Phone         string `json:"phone,omitempty" valid:"phone"`
}

type VerifyCodeEmailRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Email         string `json:"email,omitempty" valid:"email"`
}

func VerifyCodePhone(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:Phone numbers are required.",
			"digits:Must be 11 phone number.",
		},
		"captcha_id": []string{
			"required:Captcha ID are required.",
		},
		"captcha_answer": []string{
			"required:Captcha Answer are required.",
			"digits:Must be 6 captcha number.",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*VerifyCodePhoneRequest)
	errs = validate(_data.CaptchaID, rules, messages)

	return errs
}

func VerifyCodeEmail(data interface{}, _ *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":          []string{"required", "min:4", "max:30", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email address is required.",
			"min:Email address must be at least 5.",
			"max:Email address must be at least 30 characters.",
			"email:Invalid email address.",
		},
		"captcha_id": []string{
			"required:Captcha ID are required.",
		},
		"captcha_answer": []string{
			"required:Captcha Answer are required.",
			"digits:Must be 6 captcha number.",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*VerifyCodeEmailRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
