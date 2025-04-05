package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param		user	body	authcontroller.ForgotPassword.Req	true	"Request body"
// @Success	200
// @Failure	500	{string}	er.APIError
// @Failure	400	{object}	er.APIError
// @Failure	404	{object}	er.APIError
// @Router		/v1/forgotPassword [post]
// @Tags		Home
func (h *Handler) ForgotPassword(c *gin.Context) {
	type Req struct {
		Email string `json:"email" validate:"required"`
	}

	var req Req
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	err := h.Usecase.ForgotPassword(c.Request.Context(), authucase.ForgotPasswordInput{
		Email: req.Email,
	})
	if err != nil {
		er.Bind(c, erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
