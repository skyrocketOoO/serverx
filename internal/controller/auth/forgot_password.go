package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param		user	body	authucase.ForgotPasswordIn	true	"Request body"
// @Success	200
// @Failure	500	{object}	er.APIError
// @Failure	400	{object}	er.APIError
// @Failure	404	{object}	er.APIError
// @Router		/v1/forgotPassword [post]
// @Tags		Home
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
