package model

import (
	"github.com/google/uuid"
	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/gorm-plugin/lib/model"
	ope "github.com/skyrocketOoO/gorm-plugin/lib/operator"
	wh "github.com/skyrocketOoO/gorm-plugin/lib/where"
	col "github.com/skyrocketOoO/web-server-template/internal/gen/column"
	dm "github.com/skyrocketOoO/web-server-template/internal/global/domain"
	"gorm.io/gorm"
)

type User struct {
	model.Base

	Name     string `gorm:"unique"`
	Password string `gorm:"type:varchar(255)"`
}

func (u *User) Create(db *gorm.DB) error {
	var existingUser User
	if err := db.Where(wh.B(col.Users.Name, ope.Eq), u.Name).
		Take(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return erx.W(err)
		}
	} else {
		return erx.W(dm.ErrUserNameRepetite)
	}

	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	if err := db.Create(u).Error; err != nil {
		return erx.W(err)
	}
	return nil
}
