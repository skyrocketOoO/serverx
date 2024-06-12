package cmd

import (
	"os"

	restapi "web-server-template/api/rest"
	"web-server-template/internal/boot"
	"web-server-template/internal/controller/rest"
	"web-server-template/internal/controller/rest/middleware"
	"web-server-template/internal/repository/orm"
	"web-server-template/internal/usecase"
	"web-server-template/manifest/config"

	errors "github.com/rotisserie/eris"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func workFunc(cmd *cobra.Command, args []string) {
	boot.InitAll()

	if err := config.ReadConfig(); err != nil {
		log.Fatal().Msg(errors.ToString(err, true))
	}

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
