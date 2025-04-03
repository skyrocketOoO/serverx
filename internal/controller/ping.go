package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/global"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Summary Check the server started
// @Accept json
// @Produce json
// @Success 200
// @Router /ping [get]
func (d *Handler) Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}

// @Summary Check the services are healthy
// @Success 200 {object} util.ErrResp
// @Failure 503 {object} util.ErrResp
// @Router /healthy [get]
func (d *Handler) Healthy(c *gin.Context) {
	// do something check
	db := global.DB
	sqlDb, err := db.DB()
	if err != nil {
		util.RespErr(c, http.StatusServiceUnavailable, erx.W(err))
		return
	}
	if err := sqlDb.Ping(); err != nil {
		util.RespErr(c, http.StatusServiceUnavailable, erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
