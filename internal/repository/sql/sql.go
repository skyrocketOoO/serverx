package sql

import "gorm.io/gorm"

type OrmRepository struct {
}

func NewOrmRepository(db *gorm.DB) (*OrmRepository, error) {
	return &OrmRepository{}, nil
}
