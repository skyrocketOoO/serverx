package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// @Summary Check the server started
// @Accept json
// @Produce json
// @Success 200 {object} dm.ErrResp
// @Router /ping [get]
func (d *Handler) Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}

// @Summary Check the server healthy
// @Accept json
// @Produce json
// @Success 200 {object} dm.ErrResp
// @Failure 503 {object} dm.ErrResp
// @Router /healthy [get]
func (d *Handler) Healthy(c *gin.Context) {
	// do something check

	c.Status(http.StatusOK)
}
