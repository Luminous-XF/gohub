package user

import (
	"gohub/pkg/hash"

	"gorm.io/gorm"
)

func (model *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(model.Password) {
		model.Password = hash.BcryptHash(model.Password)
	}
	return
}
