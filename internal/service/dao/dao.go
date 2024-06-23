package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	Dao struct {
		db *gorm.DB
	}

	Pager struct {
		Number int
		Size   int
	}
)

func NewDao(db *gorm.DB) *Dao {
	return &Dao{db: db}
}

// Get page items with total count, src must pass the pointer of item
func (d *Dao) List(db *gorm.DB, pager Pager, src any) (total int64, err error) {
	if err = db.Count(&total).Error; err != nil {
		return
	}
	err = db.Offset(pager.Size * (pager.Number - 1)).Limit(pager.Size).Find(src).Error
	return
}

func (d *Dao) Ping(ctx context.Context) error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}
