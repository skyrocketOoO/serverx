/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"time"

	restapi "web-server-template/api/rest"
	docs "web-server-template/docs/rest"
	"web-server-template/internal/controller/rest"
	"web-server-template/internal/controller/rest/middleware"
	"web-server-template/internal/repository/orm"
	"web-server-template/internal/usecase"
	"web-server-template/manifest/config"

	errors "github.com/rotisserie/eris"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func workFunc(cmd *cobra.Command, args []string) {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(
		zerolog.ConsoleWriter{Out: os.Stderr},
	) // human-friendly logging without efficiency
	log.Info().Msg("Logger initialized")

	if err := config.ReadConfig(); err != nil {
		log.Fatal().Msg(errors.ToString(err, true))
	}

	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http"}

	db, _ := cmd.Flags().GetString("database")
	sqlDb, err := orm.InitDB(db)
	if err != nil {
		log.Fatal().Msg(errors.ToString(err, true))
	}
	defer func() {
		db, _ := sqlDb.DB()
		db.Close()
	}()

	sqlRepo, err := orm.NewOrmRepository(sqlDb)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	usecase := usecase.NewUsecase(sqlRepo)
	restDelivery := rest.NewRestDelivery(usecase)

	router := gin.Default()
	router.Use(middleware.CORS())
	restapi.Binding(router, restDelivery)

	port, _ := cmd.Flags().GetString("port")
	router.Run(":" + port)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "A brief description of your application",
	Long:  `The longer description`,
	Run:   workFunc,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.Flags().StringP("port", "p", "8080", "port")
	rootCmd.Flags().
		StringP("database", "d", "pg", `database enum. allowed: "pg", "sqlite"`)
}
