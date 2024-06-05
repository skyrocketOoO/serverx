package usecase

import (
	"context"

	"web-server-template/internal/repository/orm"

	"gorm.io/gorm"
)

type BasicUsecase struct {
	repo orm.Repository
	db   *gorm.DB
}

func NewBasicUsecase(ormRepo orm.Repository) *BasicUsecase {
	return &BasicUsecase{
		repo: ormRepo,
	}
}

func (u *BasicUsecase) Healthy(c context.Context) error {
	// do something check like db connection is established
	if err := u.repo.Ping(c); err != nil {
		return err
	}

	return nil
}
