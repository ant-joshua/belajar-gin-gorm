package repository

import (
	"belajar-go-orm/models"
	"context"
	"gorm.io/gorm"
)

type OrderRepository interface {
	InsertOrder(ctx context.Context, tx *gorm.DB, order *models.Order) (*models.Order, error)
	BeginTransaction() *gorm.DB
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func (o OrderRepositoryImpl) BeginTransaction() *gorm.DB {
	return o.db.Begin()
}

func (o OrderRepositoryImpl) InsertOrder(ctx context.Context, tx *gorm.DB, order *models.Order) (*models.Order, error) {
	err := tx.Debug().WithContext(ctx).Create(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
	}
}
