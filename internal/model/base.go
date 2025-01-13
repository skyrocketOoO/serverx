package model

import "gorm.io/gorm"

type Base struct {
	ID string `gorm:"type:char(36);default:uuid();primaryKey"`
	gorm.Model
}
