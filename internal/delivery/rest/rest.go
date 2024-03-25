package rest

import (
	"go-server-template/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestDelivery struct {
	usecase *usecase.Usecase
}

func NewRestDelivery(usecase *usecase.Usecase) *RestDelivery {
	return &RestDelivery{
		usecase: usecase,
	}
}

func (d *RestDelivery) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (d *RestDelivery) Healthy(c *gin.Context) {
	// do something check like db connection is established
	if err := d.usecase.Healthy(c.Request.Context()); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}
