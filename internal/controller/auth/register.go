package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param   user  body  auth.Register.Req  true  "Register User"
// @Success 200
// @Failure 500 {string} domain.ErrResp
// @Failure 400 {object} domain.ErrResp
// @Router /Register [post]
func (h *Handler) Register(c *gin.Context) {
	type Req struct {
		Email    string `json:"email"    validate:"required"`
		NickName string `json:"nickName" validate:"required"`
	}

	var req Req
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	err := h.Usecase.Register(c.Request.Context(), auth.RegisterInput{
		Email:    req.Email,
		NickName: req.NickName,
	})
	if err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
