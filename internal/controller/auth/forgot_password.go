package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Tags		auth
// @Param		user	body	authucase.ForgotPasswordIn	true	"Request body"
// @Success		200
// @Router		/v1/forgot-password [post]
func (h *Handler) ForgotPassword(c *gin.Context) {
	var req authucase.ForgotPasswordIn
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	if err := h.Usecase.ForgotPassword(c.Request.Context(), req); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
