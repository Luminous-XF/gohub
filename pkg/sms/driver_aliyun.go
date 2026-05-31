package sms

import (
	"encoding/json"
	"gohub/pkg/config"
	"gohub/pkg/logger"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dypnsapi "github.com/alibabacloud-go/dypnsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

type Aliyun struct{}

func (sms *Aliyun) Send(phone string, message Message) bool {
	credential, err := credentials.NewCredential(&credentials.Config{
		Type:            tea.String("access_key"),
		AccessKeyId:     tea.String(config.GetString("sms.aliyun.access_key_id")),
		AccessKeySecret: tea.String(config.GetString("sms.aliyun.access_key_secret")),
	})
	if err != nil {
		logger.ErrorString("SMS[AliYun]", "NewCredential", err.Error())
		return false
	}

	client, err := dypnsapi.NewClient(&openapi.Config{
		Credential: credential,
		Endpoint:   tea.String(config.GetString("sms.aliyun.endpoint")),
	})
	if err != nil {
		logger.ErrorString("SMS[AliYun]", "NewClient", err.Error())
		return false
	}

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("SMS[AliYun]", "Marshal Message", err.Error())
		return false
	}

	sendVerifyCodeRequest := &dypnsapi.SendSmsVerifyCodeRequest{
		SignName:      tea.String(config.GetString("sms.aliyun.sign_name")),
		TemplateCode:  tea.String(config.GetString("sms.aliyun.template_code")),
		PhoneNumber:   tea.String(phone),
		TemplateParam: tea.String(string(templateParam)),
	}
	runtime := &util.RuntimeOptions{}

	response, err := client.SendSmsVerifyCodeWithOptions(sendVerifyCodeRequest, runtime)
	if err != nil {
		logger.ErrorString("SMS[AliYun]", "SendSmsVerifyCode", err.Error())
		return false
	}

	if response.Body == nil || *response.Body.Code != "OK" {
		if response.Body != nil {
			logger.ErrorString("SMS[AliYun]", "SendSmsVerifyCode", *response.Body.Message)
		} else {
			logger.ErrorString("SMS[AliYun]", "SendSmsVerifyCode", "Response body is nil")
		}
		return false
	}

	logger.InfoString("SMS[AliYun]", "SendSmsVerifyCode", *response.Body.Code)

	return true
}
