package model

import "gorm.io/gorm"

type User struct {
	ID string `gorm:"type:char(36);default:uuid();primaryKey"`
	gorm.Model

	Name     string `gorm:"unique"`
	Password string `gorm:"type:varchar(255)"`
}
