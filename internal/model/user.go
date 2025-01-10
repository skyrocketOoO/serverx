package model

import (
	"github.com/skyrocketOoO/gorm-plugin/lib/model"
)

type User struct {
	model.Base

	Name     string `gorm:"unique"`
	Password string `gorm:"type:varchar(255)"`
}
