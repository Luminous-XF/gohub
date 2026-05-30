package sms

import (
	"encoding/json"
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/logger"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type Aliyun struct{}

func (sms *Aliyun) Send(phone string, message Message) bool {
	client, err := dysmsapi.NewClientWithAccessKey(
		"cn-hangzhou",
		config.GetString("sms.aliyun.access_key_id"),
		config.GetString("sms.aliyun.access_key_secret"),
	)
	if err != nil {
		logger.ErrorString("SMS[AliYun]", "初始化客户端失败", err.Error())
		return false
	}

	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = phone
	req.SignName = config.GetString("sms.aliyun.sign_name")
	req.TemplateCode = config.GetString("sms.aliyun.template_code")
	templateParam, _ := json.Marshal(message.Data)
	req.TemplateParam = string(templateParam)

	rsp, err := client.SendSms(req)
	if err != nil {
		logger.ErrorString("SMS[SliYun]", "Send request failed", err.Error())
		return false
	}

	if rsp.Code != "OK" {
		errMsg := fmt.Sprintf("code: %s, msg: %s, requestId: %s", rsp.Code, rsp.Message, rsp.RequestId)
		logger.ErrorString("SMS[SliYun", "Send sms failed", errMsg)
		return false
	}

	logger.InfoString(
		"SMS[AliYun]",
		"Send sms success",
		fmt.Sprintf("Phone number: %s, RequestId: %s", phone, rsp.RequestId),
	)

	return true
}
