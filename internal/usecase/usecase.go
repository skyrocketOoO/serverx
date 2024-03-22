package usecase

import (
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
