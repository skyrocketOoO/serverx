package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/skyrocketOoO/web-server-template/docs/openapi"
	"github.com/skyrocketOoO/web-server-template/internal/controller"
	"github.com/skyrocketOoO/web-server-template/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Bind(r *gin.Engine, h *controller.Handler) {
	r.GET("/ping", h.Ping)
	r.GET("/healthy", h.Healthy)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	verR := r.Group("/v1")
	pR := verR.Group("/")
	pR.Use(middleware.Jwt())
}
