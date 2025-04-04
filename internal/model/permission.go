package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model

	Name  string
	Roles []*Role `gorm:"many2many:role_permissions"`
}
