package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (model *User) Create() {
	database.DB.Create(&model)
}

func (model *User) ComparePassword(password string) bool {
	return hash.BcryptCheck(password, model.Password)
}
