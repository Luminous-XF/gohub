package config

import "gohub/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME", "速通互联验证码"),
				"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE", "SMS_335130238"),
				"template_param":    config.Env("SMS_ALIYUN_TEMPLATE_PARAM", `{"code":"##code##","min":"5"}`),
				"endpoint":          config.Env("SMS_ALIYUN_ENDPOINT", "dypnsapi.aliyuncs.com"),
			},
		}
	})
}
