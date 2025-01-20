package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/go-utils/auth"
	ope "github.com/skyrocketOoO/gorm-plugin/lib/operator"
	wh "github.com/skyrocketOoO/gorm-plugin/lib/where"
	cm "github.com/skyrocketOoO/serverx/internal/common"
	col "github.com/skyrocketOoO/serverx/internal/gen/column"
	"github.com/skyrocketOoO/serverx/internal/global"
	dm "github.com/skyrocketOoO/serverx/internal/global/domain"
	"github.com/skyrocketOoO/serverx/internal/model"
	"gorm.io/gorm"
)

// @Param request body controller.CreateUser.Req true "Request body"
// @Failure 400 {object} dm.ErrResp ""
// @Success 200
// @Failure 500 {object} dm.ErrResp ""
// @Router /user/create [post]
// @Security Bearer
// @Tags Alarm
func (d *Handler) CreateUser(c *gin.Context) {
	type Req struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var req Req
	if ok := cm.BindAndValidate(c, &req); !ok {
		return
	}

	db := global.DB
	var existingUser model.User
	if err := db.Where(wh.B(col.Users.Name, ope.Eq), req.Name).
		Take(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			dm.RespErr(c, dm.ToHttpCode(err), erx.W(err))
			return
		}
	} else {
		err = dm.ErrUserNameRepetite
		dm.RespErr(c, dm.ToHttpCode(err), erx.W(dm.ErrUserNameRepetite))
		return
	}

	if err := db.Create(&model.User{
		Name:     req.Name,
		Password: string(auth.Hash(req.Password, cm.GetSalt())),
	}).Error; err != nil {
		dm.RespErr(c, dm.ToHttpCode(err), erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}

// @Param request body controller.GetUsers.Req true "Request body"
// @Failure 400 {object} dm.ErrResp ""
// @Success 200 {object} controller.GetUsers.Resp ""
// @Failure 500 {object} dm.ErrResp ""
// @Router /user/get [post]
// @Security Bearer
// @Tags Alarm
func (d *Handler) GetUsers(c *gin.Context) {
	type Req struct {
		Pager  *cm.Pager   `json:"pager"`
		Sorter []cm.Sorter `json:"sorter"`
	}

	var req Req
	if ok := cm.BindAndValidate(c, &req); !ok {
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
	if err := db.Model(&model.User{}).Count(&resp.Count).Error; err != nil {
		dm.RespErr(c, dm.ToHttpCode(err), erx.W(err))
		return
	}

	if err := db.Model(&model.User{}).Scan(&resp.Data).Error; err != nil {
		dm.RespErr(c, dm.ToHttpCode(err), erx.W(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Param request body controller.UpdateUser.Req true "Request body"
// @Failure 400 {object} dm.ErrResp ""
// @Success 200
// @Failure 500 {object} dm.ErrResp ""
// @Router /user/update [post]
// @Security Bearer
// @Tags Alarm
func (d *Handler) UpdateUser(c *gin.Context) {
	type Req struct {
		ID   uint   `json:"id" validate:"required"`
		Name string `json:"name"`
	}

	var req Req
	if ok := cm.BindAndValidate(c, &req); !ok {
		return
	}

	db := global.DB

	// Find the user
	var user model.User
	if err := db.Take(&user, req.ID).Error; err != nil {
		dm.RespErr(c, dm.ToHttpCode(err), erx.W(err))
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if err := db.Save(&user).Error; err != nil {
		dm.RespErr(c, dm.ToHttpCode(err), erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}

// @Param request body controller.DeleteUser.Req true "Request body"
// @Failure 400 {object} dm.ErrResp ""
// @Success 200
// @Failure 500 {object} dm.ErrResp ""
// @Router /user/delete [post]
// @Security Bearer
// @Tags Alarm
func (d *Handler) DeleteUser(c *gin.Context) {
	type Req struct {
		ID uint `json:"id" validate:"required"`
	}

	var req Req
	if ok := cm.BindAndValidate(c, &req); !ok {
		return
	}

	db := global.DB

	if err := db.Delete(&model.User{}, req.ID).Error; err != nil {
		dm.RespErr(c, dm.ToHttpCode(err), erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
