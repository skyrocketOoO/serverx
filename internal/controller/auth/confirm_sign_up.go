package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Tags			auth
// @Param			user	body	authucase.ConfirmSignUpIn	true	"Request body"
// @Success		200
// @Failure		500	{object}	er.APIError
// @Failure		400	{object}	er.APIError
// @Router			/v1/confirm-sign-up [post]
func (h *Handler) ConfirmSignUp(c *gin.Context) {
	var req authucase.ConfirmSignUpIn
	if !util.ParseValidate(c, &req) {
		return
	}

	if err := h.Usecase.ConfirmSignUp(c.Request.Context(), req); err != nil {
		er.Bind(c, erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
