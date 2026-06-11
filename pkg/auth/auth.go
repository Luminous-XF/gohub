package auth

import (
	"errors"
	"gohub/app/models/user"
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
