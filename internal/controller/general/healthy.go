package general

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Success 200 {object} domain.ErrResp
// @Failure 503 {object} domain.ErrResp
// @Router /healthy [get]
func (d *Handler) Healthy(c *gin.Context) {
	if err := d.usecase.Healthy(c); err != nil {
		util.RespErr(c, http.StatusServiceUnavailable, erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
