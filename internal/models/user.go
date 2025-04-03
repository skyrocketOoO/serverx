package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"unique"`
	Password string `gorm:"type:varchar(255)"`

	RoleID string `gorm:"type:char(36)"`
	Role   Role   `gorm:"foreignKey:RoleID"`
}
