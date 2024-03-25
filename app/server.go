package main

import (
	"go-server-template/api"
	"go-server-template/config"
	"go-server-template/internal/delivery/rest"
	"go-server-template/internal/delivery/rest/middleware"
	"go-server-template/internal/repository/sql"
	"go-server-template/internal/usecase"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}) // human-friendly logging without efficiency
	log.Info().Msg("Logger initialized")

	if err := config.ReadConfig(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	sqlDb, err := sql.InitDB()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	sqlRepo, err := sql.NewOrmRepository(sqlDb)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	usecase := usecase.NewUsecase(sqlRepo)
	restDelivery := rest.NewRestDelivery(usecase)

	router := gin.Default()
	router.Use(middleware.CORS())
	api.Binding(router, restDelivery)

	router.Run(viper.GetString("address"))
}
