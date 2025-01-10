package model

import (
	"github.com/skyrocketOoO/gorm-plugin/lib/model"
)

type User struct {
	model.Base

	Name     string `gorm:"unique; comment:登入帳號"`
	Password string `gorm:"type:varchar(255); comment:登入密碼"`
}
