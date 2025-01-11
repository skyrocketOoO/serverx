package global

import "gorm.io/gorm"

var (
	// env | flag
	Database string

	// instance
	DB *gorm.DB
)
