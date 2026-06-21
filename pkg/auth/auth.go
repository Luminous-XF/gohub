package auth

import (
	"errors"
	"gohub/app/models/user"
	"gohub/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Attempt(email, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("user not found")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("password wrong")
	}

	return userModel, nil
}

func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("user not found")
	}

	return userModel, nil
}

func CurrentUser(ctx *gin.Context) user.User {
	userModel, ok := ctx.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("failed to retrieve user"))
		return user.User{}
	}

	return userModel
}

func CurrentUID(ctx *gin.Context) string {
	return ctx.GetString("current_user_id")
}
