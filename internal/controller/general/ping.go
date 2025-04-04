package generalcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Success		200
// @Router			/ping [get]
// @Tags			TroubleShooting
func (h *Handler) Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}
