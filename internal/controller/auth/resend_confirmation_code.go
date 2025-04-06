package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Tags	auth
// @Param   user  body  authucase.ResendConfirmationCodeIn  true "request body"
// @Success 200
// @Router /v1/resend-confirmation-code [post]
func (h *Handler) ResendConfirmationCode(c *gin.Context) {
	var req authucase.ResendConfirmationCodeIn
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	if err := h.Usecase.ResendConfirmationCode(c.Request.Context(), req); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
