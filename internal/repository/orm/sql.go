package orm

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Ping(c context.Context) error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	if err := db.PingContext(c); err != nil {
		return err
	}

	return nil
}
