package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Tags			auth
// @Param   user  body  authucase.RefreshTokenIn  true  "req"
// @Success 200 {object} authucase.RefreshTokenOut
// @Failure 500 {object} er.APIError
// @Failure 400 {object} er.APIError
// @Router /v1/refresh-token [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	var req authucase.RefreshTokenIn
	if !util.ParseValidate(c, &req) {
		return
	}

	resp, err := h.Usecase.RefreshToken(c.Request.Context(), req)
	if err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
