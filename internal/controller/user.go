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

func CreateUser(c *gin.Context) {
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
			dm.RespErr(c, http.StatusInternalServerError, erx.W(err))
			return
		}
	} else {
		dm.RespErr(c, http.StatusInternalServerError, erx.W(dm.ErrUserNameRepetite))
		return
	}

	if err := db.Create(&model.User{
		Name:     req.Name,
		Password: string(auth.Hash(req.Password, cm.GetSalt())),
	}).Error; err != nil {
		dm.RespErr(c, http.StatusInternalServerError, erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}

func GetUsers(c *gin.Context) {
	db := global.DB

	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		dm.RespErr(c, http.StatusInternalServerError, erx.W(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func UpdateUser(c *gin.Context) {
	type Req struct {
		ID       uint   `json:"id" validate:"required"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var req Req
	if ok := cm.BindAndValidate(c, &req); !ok {
		return
	}

	db := global.DB

	// Find the user
	var user model.User
	if err := db.First(&user, req.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			dm.RespErr(c, http.StatusNotFound, erx.W(err))
			return
		}
		dm.RespErr(c, http.StatusInternalServerError, erx.W(err))
		return
	}

	// Update fields if provided
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Password != "" {
		user.Password = string(auth.Hash(req.Password, cm.GetSalt()))
	}

	if err := db.Save(&user).Error; err != nil {
		dm.RespErr(c, http.StatusInternalServerError, erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}

func DeleteUser(c *gin.Context) {
	type Req struct {
		ID uint `json:"id" validate:"required"`
	}

	var req Req
	if ok := cm.BindAndValidate(c, &req); !ok {
		return
	}

	db := global.DB

	// Delete the user
	if err := db.Delete(&model.User{}, req.ID).Error; err != nil {
		dm.RespErr(c, http.StatusInternalServerError, erx.W(err))
		return
	}

	c.Status(http.StatusOK)
}
