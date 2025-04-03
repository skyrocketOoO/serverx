package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/go-utils/auth"
	ope "github.com/skyrocketOoO/gorm-plugin/lib/operator"
	wh "github.com/skyrocketOoO/gorm-plugin/lib/where"
	col "github.com/skyrocketOoO/serverx/internal/gen/column"
	"github.com/skyrocketOoO/serverx/internal/global"
	models "github.com/skyrocketOoO/serverx/internal/model"
	"github.com/skyrocketOoO/serverx/internal/util"
	"gorm.io/gorm"
)

// @Param request body controller.CreateUser.Req true "Request body"
// @Failure 400 {object} util.ErrResp ""
// @Success 200
// @Failure 500 {object} util.ErrResp ""
// @Router /user/create [post]
// @Security Bearer
// @Tags Alarm
func (d *Handler) CreateUser(c *gin.Context) {
	type Req struct {
		Name     string `json:"name"     validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var req Req
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	db := global.DB
	var existingUser models.User
	if err := db.Where(wh.B(col.Users.Name, ope.Eq), req.Name).
		Take(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			util.RespErr(c, util.ToHttpCode(err), erx.W(err))
			return
		}
	} else {
		err = util.ErrUserNameRepetite
		util.RespErr(c, util.ToHttpCode(err), erx.W(util.ErrUserNameRepetite))
		return
	}

	if err := db.Create(&models.User{
		Name:     req.Name,
		Password: string(auth.Hash(req.Password, util.GetSalt())),
	}).Error; err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}

// @Param request body controller.GetUsers.Req true "Request body"
// @Failure 400 {object} util.ErrResp ""
// @Success 200 {object} controller.GetUsers.Resp ""
// @Failure 500 {object} util.ErrResp ""
// @Router /user/get [post]
// @Security Bearer
// @Tags Alarm
func (d *Handler) GetUsers(c *gin.Context) {
	type Req struct {
		Pager  *util.Pager   `json:"pager"`
		Sorter []util.Sorter `json:"sorter"`
	}

	var req Req
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	db := global.DB

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	type Resp struct {
		Data  []User `json:"data"`
		Count int64  `json:"count"`
	}

	var resp Resp
	if err := db.Model(&models.User{}).Count(&resp.Count).Error; err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	if err := db.Model(&models.User{}).Scan(&resp.Data).Error; err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Param request body controller.UpdateUser.Req true "Request body"
// @Failure 400 {object} util.ErrResp ""
// @Success 200
// @Failure 500 {object} util.ErrResp ""
// @Router /user/update [post]
// @Security Bearer
// @Tags Alarm
func (d *Handler) UpdateUser(c *gin.Context) {
	type Req struct {
		ID   uint   `json:"id"   validate:"required"`
		Name string `json:"name"`
	}

	var req Req
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	db := global.DB

	// Find the user
	var user models.User
	if err := db.Take(&user, req.ID).Error; err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if err := db.Save(&user).Error; err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}

// @Param request body controller.DeleteUser.Req true "Request body"
// @Failure 400 {object} util.ErrResp ""
// @Success 200
// @Failure 500 {object} util.ErrResp ""
// @Router /user/delete [post]
// @Security Bearer
// @Tags Alarm
func (d *Handler) DeleteUser(c *gin.Context) {
	type Req struct {
		ID uint `json:"id" validate:"required"`
	}

	var req Req
	if ok := util.ParseValidate(c, &req); !ok {
		return
	}

	db := global.DB

	if err := db.Delete(&models.User{}, req.ID).Error; err != nil {
		util.RespErr(c, util.ToHttpCode(err), erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
