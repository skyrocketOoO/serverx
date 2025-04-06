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
// @Router /v1/invite-user [post]
func (h *Handler) InviteUser(c *gin.Context) {
	var req authucase.InviteUserIn
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	if err := h.Usecase.InviteUser(c.Request.Context(), req); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
