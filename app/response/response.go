package response

import (
	"errors"
	"gohub/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func Success(ctx *gin.Context) {
	JSON(ctx, gin.H{
		"success": true,
		"message": "Success!",
	})
}

func Data(ctx *gin.Context, data interface{}) {
	JSON(ctx, gin.H{
		"success": true,
		"data":    data,
	})
}

func Created(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

func CreateJSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, data)
}

func Abort404(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("Data Not Found!", msg...),
	})
}

func Abort403(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("Data Not Allowed!", msg...),
	})
}

func Abort500(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("Something went wrong, please try again later.", msg...),
	})
}

func BadRequest(ctx *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("Bad Request!", msg...),
		"error":   err.Error(),
	})
}

func Error(ctx *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		Abort404(ctx)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("Unprocessable Entity", msg...),
		"error":   err.Error(),
	})
}

func ValidationError(ctx *gin.Context, errors map[string][]string) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("Request validation failed!"),
		"error":   errors,
	})
}

func Unauthorized(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("Unauthorized!", msg...),
	})
}

func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
