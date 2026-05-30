package requests

import (
	"gohub/app/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(ctx *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 解析请求
	if err := ctx.ShouldBind(obj); err != nil {
		response.BadRequest(ctx, err, "Parse request failed.")
		return false
	}

	// 验证表单
	errs := handler(obj, ctx)
	if len(errs) > 0 {
		response.ValidationError(ctx, errs)
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, msg govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      msg,
	}

	return govalidator.New(opts).ValidateStruct()
}
