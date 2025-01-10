package server

import (
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/api"
	"github.com/skyrocketOoO/serverx/internal/boot"
	"github.com/skyrocketOoO/serverx/internal/controller"
	"github.com/skyrocketOoO/serverx/internal/controller/middleware"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "The main service command",
	Long:  ``,
	// Args:  cobra.MinimumNArgs(1),
	Run: RunServer,
}

func RunServer(cmd *cobra.Command, args []string) {
	if err := boot.InitAll(); err != nil {
		log.Fatal().Msgf("Initialization failed: %v", err)
	}

	// dbConf, _ := cmd.Flags().GetString("database")
	// err := db.New(dbConf)
	// if err != nil {
	// 	log.Fatal().Msgf("Initialization failed: %v", err)
	// }

	restController := controller.NewHandler()

	router := gin.Default()
	router.Use(middleware.Cors())
	api.Bind(router, restController)

	port, _ := cmd.Flags().GetString("port")
	router.Run(":" + port)
}

func init() {
	Cmd.Flags().StringP("port", "p", "8080", "port")
	Cmd.Flags().
		StringP("database", "d", "postgres", `database enum. allowed: "postgres", "sqlite"`)
}
