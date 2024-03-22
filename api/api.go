package api

import (
	"go-server-template/internal/delivery/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Binding(r *gin.Engine, d *rest.RestDelivery) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})
}
