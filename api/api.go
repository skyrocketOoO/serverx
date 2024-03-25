package api

import (
	"go-server-template/internal/delivery/rest"

	"github.com/gin-gonic/gin"
)

func Binding(r *gin.Engine, d *rest.RestDelivery) {
	r.GET("/ping", d.Ping)
	r.GET("/healthy", d.Healthy)
}
