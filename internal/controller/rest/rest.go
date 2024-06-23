package rest

import (
	"net/http"

	"web-server-template/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RestController struct {
	usecase *usecase.BasicUsecase
}

func NewRestController(usecase *usecase.BasicUsecase) *RestController {
	return &RestController{
		usecase: usecase,
	}
}

// @Summary Check the server started
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response
// @Router /ping [get]
func (d *RestController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "pong"})
}

// @Summary Check the server healthy
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 503 {object} domain.Response
// @Router /healthy [get]
func (d *RestController) Healthy(c *gin.Context) {
	// do something check
	if err := d.usecase.Healthy(c.Request.Context()); err != nil {
		c.JSON(http.StatusServiceUnavailable, Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "healthy"})
}
