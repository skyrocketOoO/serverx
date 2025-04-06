package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Tags	auth
// @Param   user  body  authucase.LoginIn  true "request body"
// @Success 200 {object} authucase.LoginOut
// @Router /v1/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req authucase.LoginIn
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	out, err := h.Usecase.Login(c.Request.Context(), req)
	if err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.JSON(http.StatusOK, out)
}
