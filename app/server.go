package main

import (
	"go-server-template/api"
	"go-server-template/config"
	"go-server-template/internal/delivery/rest"
	"go-server-template/internal/delivery/rest/middleware"
	"go-server-template/internal/repository/sql"
	"go-server-template/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Info().Msg("hello world")

	if err := config.ReadConfig(); err != nil {
		panic(err.Error())
	}

	sqlDb, err := sql.InitDB()
	if err != nil {
		panic(err.Error())
	}

	sqlRepo, err := sql.NewOrmRepository(sqlDb)
	if err != nil {
		panic(err.Error())
	}

	usecase := usecase.NewUsecase(sqlRepo)
	restDelivery := rest.NewRestDelivery(usecase)

	r := gin.Default()
	r.Use(middleware.CORS())
	api.Binding(r, restDelivery)

	r.Run(viper.GetString("address"))
}
