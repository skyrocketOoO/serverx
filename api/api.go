package api

import (
	"github.com/gin-gonic/gin"
	controller "github.com/skyrocketOoO/serverx/internal/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterAPIHandlers(r *gin.Engine, h *controller.Handler) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", h.Ping)
	r.GET("/healthy", h.General.Healthy)

	r.POST("/login", h.Auth.Login)
	r.POST("/register", h.Auth.Register)

	// pR := r.Group("/")
	// pR.Use(middleware.Jwt())
}
