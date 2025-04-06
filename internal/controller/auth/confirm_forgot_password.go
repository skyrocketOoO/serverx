package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param	user	body	authucase.ConfirmForgotPasswordIn	true	"Request body"
// @Success	200
// @Router	/v1/confirm-forgot-password [post]
// @Tags	auth
func (h *Handler) ConfirmForgotPassword(c *gin.Context) {
	var req authucase.ConfirmForgotPasswordIn
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	if err := h.Usecase.ConfirmForgotPassword(c.Request.Context(), req); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
