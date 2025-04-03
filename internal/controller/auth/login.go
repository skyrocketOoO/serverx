package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param   user  body  auth.Login.Req  true  "Login User"
// @Success 200 {object} auth.Login.Resp "token"
// @Failure 500 {string} util.ErrResp
// @Failure 400 {object} util.ErrResp
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	type Req struct {
		Name     string `json:"Name"     validate:"required"`
		Password string `json:"Password" validate:"required"`
	}

	var req Req
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	token, err := h.Usecase.Login(c.Request.Context(), auth.LoginInput{
		Email:    req.Name,
		Password: req.Password,
	})
	if err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	type Resp struct {
		Token string `json:"token"`
	}

	c.JSON(http.StatusOK, Resp{Token: token})
}
