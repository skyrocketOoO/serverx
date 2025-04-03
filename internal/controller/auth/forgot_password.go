package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param		user	body	controller.ForgotPassword.Req	true	"Request body"
// @Success	200
// @Failure	500	{string}	dom.ErrResp
// @Failure	400	{object}	dom.ErrResp
// @Failure	404	{object}	dom.ErrResp
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

	err := h.Usecase.ForgotPassword(c.Request.Context(), auth.ForgotPasswordInput{
		Email: req.Email,
	})
	if err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
