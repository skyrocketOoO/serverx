package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param   user  body  authcontroller.Register.Req  true  "Register User"
// @Success 200
// @Failure 500 {string} er.APIError
// @Failure 400 {object} er.APIError
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

	err := h.Usecase.Register(c.Request.Context(), authucase.RegisterInput{
		Email:    req.Email,
		NickName: req.NickName,
	})
	if err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.Status(http.StatusOK)
}
