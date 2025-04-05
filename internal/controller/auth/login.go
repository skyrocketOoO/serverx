package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	"github.com/skyrocketOoO/serverx/internal/util"
)

// @Param   user  body  authcontroller.Login.Req  true  "Login User"
// @Success 200 {object} authucase.LoginOut
// @Failure 500 {string} er.APIError
// @Failure 400 {object} er.APIError
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

	out, err := h.Usecase.Login(c.Request.Context(), authucase.LoginIn{
		Email:    req.Name,
		Password: req.Password,
	})
	if err != nil {
		er.Bind(c, er.W(err))
		return
	}

	c.JSON(http.StatusOK, out)
}
