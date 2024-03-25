package rest

import (
	"go-server-template/domain"
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
	c.JSON(http.StatusOK, domain.Response{Message: "pong"})
}

func (d *RestDelivery) Healthy(c *gin.Context) {
	// do something check
	if err := d.usecase.Healthy(c.Request.Context()); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "healthy"})
}
