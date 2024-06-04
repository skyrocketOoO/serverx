package usecase

import (
	"context"

	"web-server-template/internal/repository/orm/domain"
)

type BasicUsecase struct {
	repo domain.OrmRepository
}

func NewUsecase(ormRepo domain.OrmRepository) *BasicUsecase {
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
