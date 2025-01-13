package global

import "gorm.io/gorm"

const (
	ApiVersion = "/v1"
)

var (
	// env | flag
	Database    string
	AutoMigrate bool = false
	Env         string

	// instance
	DB *gorm.DB
)
