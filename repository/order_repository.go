package repository

import (
	"belajar-go-orm/models"
	"context"
	"gorm.io/gorm"
)

type OrderRepository interface {
	InsertOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	BeginTransaction(ctx context.Context) *gorm.DB
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func (o OrderRepositoryImpl) BeginTransaction(ctx context.Context) *gorm.DB {
	return o.db.WithContext(ctx).Begin()
}

func (o OrderRepositoryImpl) InsertOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	err := o.db.Debug().WithContext(ctx).Create(&order).Error
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
