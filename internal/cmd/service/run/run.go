package run

import (
	restapi "web-server-template/api/rest"
	"web-server-template/internal/boot"
	"web-server-template/internal/controller/rest"
	"web-server-template/internal/controller/rest/middleware"
	"web-server-template/internal/service/dao"
	"web-server-template/internal/service/orm"
	"web-server-template/internal/usecase"

	errors "github.com/rotisserie/eris"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run service",
	Long:  ``,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := boot.InitAll(); err != nil {
			panic(err.Error())
		}

		dbConf, _ := cmd.Flags().GetString("database")
		db, err := orm.NewDB(dbConf)
		if err != nil {
			log.Fatal().Msg(errors.ToString(err, true))
		}
		defer func() {
			db, _ := db.DB()
			db.Close()
		}()

		dao := dao.NewDao(db)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		usecase := usecase.NewBasicUsecase(dao)
		restDelivery := rest.NewRestDelivery(usecase)

		router := gin.Default()
		router.Use(middleware.CORS())
		restapi.Binding(router, restDelivery)

		port, _ := cmd.Flags().GetString("port")
		router.Run(":" + port)
	},
}

func init() {
	RunCmd.Flags().StringP("port", "p", "8080", "port")
	RunCmd.Flags().
		StringP("database", "d", "pg", `database enum. allowed: "pg", "sqlite"`)
}
