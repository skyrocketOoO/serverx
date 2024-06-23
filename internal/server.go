package internal

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
	restController := rest.NewRestController(usecase)

	router := gin.Default()
	router.Use(middleware.CORS())
	restapi.Binding(router, restController)

	port, _ := cmd.Flags().GetString("port")
	router.Run(":" + port)
}
