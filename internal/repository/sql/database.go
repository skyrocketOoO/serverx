package sql

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() (*gorm.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.db"),
		viper.GetString("postgres.sslmode"),
		viper.GetString("postgres.timezone"),
	)
	return gorm.Open(
		postgres.Open(connStr), &gorm.Config{
			Logger: nil,
		},
	)
}
