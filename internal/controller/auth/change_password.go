package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Tags	auth
// @Param   user  body  authcontroller.ChangePassword.Req  true  "req"
// @Success 200
// @Router /v1/change-password [post]
func (h *Handler) ChangePassword(c *gin.Context) {
	type Req struct {
		OldPass string `json:"oldPass" validate:"required"`
		NewPass string `json:"newPass" validate:"required"`
	}

	var req Req
	if !util.ParseValidate(c, &req) {
		return
	}

	if err := h.Usecase.ChangePassword(c.Request.Context(), authucase.ChangePasswordIn{
		OldPass:     req.OldPass,
		NewPass:     req.NewPass,
		AccessToken: c.GetHeader("Authorization"),
	}); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
