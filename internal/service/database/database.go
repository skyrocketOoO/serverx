package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/serverx/internal/domain"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var initOnce sync.Once

type zerologWriter struct{}

func (z *zerologWriter) Printf(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func New() (db *gorm.DB, err error) {
	initOnce.Do(func() {
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

		switch domain.Database {
		case "mysql":
			log.Info().Msg("Connecting to MySQL")
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
		}

		if err != nil {
			err = er.W(err, er.DBUnavailable)
			return
		}
	})
	return db, err
}

func Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
