package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/go-utils/auth"
	ope "github.com/skyrocketOoO/gorm-plugin/lib/operator"
	wh "github.com/skyrocketOoO/gorm-plugin/lib/where"
	cm "github.com/skyrocketOoO/web-server-template/internal/common"
	col "github.com/skyrocketOoO/web-server-template/internal/gen/column"
	tb "github.com/skyrocketOoO/web-server-template/internal/gen/table"
	dm "github.com/skyrocketOoO/web-server-template/internal/global/domain"
	"github.com/skyrocketOoO/web-server-template/internal/model"
	"github.com/skyrocketOoO/web-server-template/internal/service/exter/db"
)

// @Param   user  body  controller.Login.Req  true  "Login User"
// @Success 200 {object} controller.Login.Resp "token"
// @Failure 500 {string} dm.ErrResp "error"
// @Failure 400 {object} dm.ErrResp "bad request"
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	type Req struct {
		Name     string `json:"Name" validate:"required"`
		Password string `json:"Password" validate:"required"`
	}

	var req Req
	if ok := cm.BindAndValidate(c, &req); !ok {
		return
	}

	hashedPassword := string(auth.Hash(req.Password, cm.GetSalt()))
	db := db.Get()

	var user model.User
	if err := db.
		Table(tb.Users).
		Where(wh.B(col.Users.Name, ope.Eq), req.Name).
		Where(wh.B(col.Users.Password, ope.Eq), hashedPassword).
		Take(&user).
		Error; err != nil {
		dm.RespErr(c, http.StatusInternalServerError, erx.W(dm.ErrLoginFailed))
		return
	}

	token, err := cm.GenerateToken(user.ID)
	if err != nil {
		dm.RespErr(c, http.StatusInternalServerError, erx.W(err))
		return
	}

	token = "Bearer " + token

	type Resp struct {
		Token string `json:"token"`
	}
	c.JSON(http.StatusOK, Resp{Token: token})
}

// @Param   user  body  controller.Register.Req  true  "Login User"
// @Success 200 {object} controller.Register.Resp "token"
// @Failure 500 {string} dm.ErrResp "error"
// @Failure 400 {object} dm.ErrResp "bad request"
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	type Req struct {
		Name     string `json:"Name" validate:"required,min=6,max=32"`
		Password string `json:"Password" validate:"required,min=8,max=32"`
	}

	var req Req
	if ok := cm.BindAndValidate(c, &req); !ok {
		return
	}

	db := db.Get()
	user := model.User{
		Name:     req.Name,
		Password: string(auth.Hash(req.Password, cm.GetSalt())),
	}

	if err := user.Create(db); err != nil {
		dm.RespErr(c, http.StatusInternalServerError, erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Logout(c *gin.Context) {
}

func (h *Handler) ForgetPassword(c *gin.Context) {
}
