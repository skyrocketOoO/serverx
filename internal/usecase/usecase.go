package usecase

import (
	"context"
	"go-server-template/internal/repository/sql"
)

type Usecase struct {
	repo *sql.OrmRepository
}

func NewUsecase(ormRepo *sql.OrmRepository) *Usecase {
	return &Usecase{
		repo: ormRepo,
	}
}

func (u *Usecase) Healthy(c context.Context) error {
	// do something check like db connection is established
	if err := u.repo.Ping(c); err != nil {
		return err
	}

	return nil
}
