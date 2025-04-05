package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
// @Success 200 {object} er.APIError
// @Failure 503 {object} er.APIError
// @Router /healthy [get]
// func (d *Handler) Healthy(c *gin.Context) {
// 	// do something check
// 	db := postgres.Get()
// 	sqlDb, err := db.DB()
// 	if err != nil {
// 		util.ErrResp(c, http.StatusServiceUnavailable, erx.W(err))
// 		return
// 	}
// 	if err := sqlDb.Ping(); err != nil {
// 		util.ErrResp(c, http.StatusServiceUnavailable, erx.W(err))
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }
