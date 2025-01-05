package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/web-server-template/internal/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Binding(r *gin.Engine, d *controller.RestController) {
	r.GET("/ping", d.Ping)
	r.GET("/healthy", d.Healthy)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
