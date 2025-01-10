package db

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db      *gorm.DB
	Migrate bool = false
)

func Get() *gorm.DB {
	return db
}

type zerologWriter struct{}

func (z *zerologWriter) Printf(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func New(database string) error {
	log.Info().Msg("New db")

	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
		Logger: logger.New(
			&zerologWriter{},
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      true,
				Colorful:                  true,
			},
		),
	}
	var err error
	switch database {
	case "mysql":
		log.Info().Msg("Connecting to Postgres")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			viper.GetString("db.user"),
			viper.GetString("db.password"),
			viper.GetString("db.host"),
			viper.GetInt("db.port"),
			viper.GetString("db.db"),
			viper.GetString("db.timezone"),
		)

		db, err = gorm.Open(mysql.Open(dsn), &config)
	case "postgres":
		log.Info().Msg("Connecting to Postgres")
		connStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s",
			viper.GetString("db.host"),
			viper.GetString("db.port"),
			viper.GetString("db.user"),
			viper.GetString("db.password"),
			viper.GetString("db.db"),
			viper.GetString("db.timezone"),
		)
		db, err = gorm.Open(postgres.Open(connStr), &config)

	case "sqlite":
		log.Info().Msg("Connecting to Sqlite")
		db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	}

	if err != nil {
		return erx.W(err)
	}

	if Migrate {
		if err = db.AutoMigrate(
			&model.User{},
		); err != nil {
			return err
		}
	}
	return nil
}
