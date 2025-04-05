package generalcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
)

// @Success 200 {object} er.APIError
// @Failure 503 {object} er.APIError
// @Router /healthy [get]
func (d *Handler) Healthy(c *gin.Context) {
	if err := d.usecase.Healthy(c); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
