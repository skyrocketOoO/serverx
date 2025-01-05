package internal

import (
	errors "github.com/rotisserie/eris"
	"github.com/skyrocketOoO/web-server-template/api"
	"github.com/skyrocketOoO/web-server-template/internal/boot"
	"github.com/skyrocketOoO/web-server-template/internal/controller"
	"github.com/skyrocketOoO/web-server-template/internal/middleware"
	"github.com/skyrocketOoO/web-server-template/internal/service/dao"
	"github.com/skyrocketOoO/web-server-template/internal/service/orm"
	"github.com/skyrocketOoO/web-server-template/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func RunServer(cmd *cobra.Command, args []string) {
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
	restController := controller.NewRestController(usecase)

	router := gin.Default()
	router.Use(middleware.CORS())
	api.Bind(router, restController)

	port, _ := cmd.Flags().GetString("port")
	router.Run(":" + port)
}
