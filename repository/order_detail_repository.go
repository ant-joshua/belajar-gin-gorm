package repository

import (
	"belajar-go-orm/models"
	"context"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	InsertOrderDetail(ctx context.Context, order *models.OrderDetail) (*models.OrderDetail, error)
}

type OrderDetailRepositoryImpl struct {
	db *gorm.DB
}

func (o OrderDetailRepositoryImpl) InsertOrderDetail(ctx context.Context, orderDetail *models.OrderDetail) (*models.OrderDetail, error) {
	err := o.db.Debug().WithContext(ctx).Create(&orderDetail).Error
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
