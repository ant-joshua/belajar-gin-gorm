package repository

import (
	"belajar-go-orm/models"
	"context"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	InsertOrderDetail(ctx context.Context, tx *gorm.DB, order *models.OrderDetail) (*models.OrderDetail, error)
}

type OrderDetailRepositoryImpl struct {
	db *gorm.DB
}

func (o OrderDetailRepositoryImpl) InsertOrderDetail(ctx context.Context, tx *gorm.DB, orderDetail *models.OrderDetail) (*models.OrderDetail, error) {
	err := tx.Debug().WithContext(ctx).Create(&orderDetail).Error
	if err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &OrderDetailRepositoryImpl{
		db: db,
	}
}
