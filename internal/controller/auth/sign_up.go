package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Tags	auth
// @Param   user  body  authucase.SignUpIn  true  "req"
// @Success 200
// @Router /v1/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var req authucase.SignUpIn
	if !util.ParseValidate(c, &req) {
		return
	}

	if err := h.Usecase.SignUp(c.Request.Context(), req); err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
