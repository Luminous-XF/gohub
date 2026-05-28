package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(ctx *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 解析请求
	if err := ctx.ShouldBind(obj); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse request failed.",
			"error":   err.Error(),
		})

		fmt.Println(err.Error())

		return false
	}

	// 验证表单
	errs := handler(obj, ctx)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation failed.",
			"error":   errs,
		})
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
