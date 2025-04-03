package api

import (
	"github.com/gin-gonic/gin"
	controller "github.com/skyrocketOoO/serverx/internal/controller"
	middleware "github.com/skyrocketOoO/serverx/internal/controller/middleware"
	"github.com/skyrocketOoO/serverx/internal/global"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterAPIHandlers(r *gin.Engine, h *controller.Handler) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	vr := r.Group("/" + global.ApiVersion)
	vr.GET("/ping", h.Ping)
	vr.GET("/healthy", h.Healthy)

	vr.POST("/login", h.Login)
	vr.POST("/register", h.Register)

	pR := vr.Group("/")
	pR.Use(middleware.Jwt())

	userR := pR.Group("/user")
	{
		userR.POST("/get", h.GetUsers)
		userR.POST("/create", h.CreateUser)
		userR.POST("/update", h.UpdateUser)
		userR.POST("/delete", h.DeleteUser)
	}
}
