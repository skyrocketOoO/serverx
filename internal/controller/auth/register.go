package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param   user  body  authcontroller.SignUp.Req  true  "Register User"
// @Success 200
// @Failure 500 {string} er.APIError
// @Failure 400 {object} er.APIError
// @Router /sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	type Req struct {
		Email    string `json:"email"    validate:"required"`
		Password string `json:"password" validate:"required"`
		NickName string `json:"nickName" validate:"required"`
	}

	var req Req
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	err := h.Usecase.SignUp(c.Request.Context(), authucase.SignUpInput{
		Email:    req.Email,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
