package models

import (
	"time"
)

type Order struct {
	ID          uint          `json:"id" gorm:"primary_key"`
	UserID      uint          `json:"user_id" gorm:"not null"`
	User        User          `json:"user" gorm:"foreignKey:UserID"`
	OrderCode   string        `json:"order_code" gorm:"type:varchar(100);not null"`
	OrderDetail []OrderDetail `json:"order_detail" gorm:"foreignKey:OrderID"`
	TotalPrice  float64       `json:"total_price" gorm:"not null;default:0.0"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at"`
}

type CreateOrderRequest struct {
	UserID  uint                       `json:"user_id" binding:"required"`
	Details []CreateOrderDetailRequest `json:"details" binding:"required"`
}

type CreateOrderDetailRequest struct {
	ProductID uint `json:"product_id" binding:"required" default:"1"`
	Quantity  int  `json:"quantity" binding:"required" default:"1"`
	Price     int  `json:"price" binding:"required" default:"10000"`
}

type OrderDetail struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	OrderID   uint       `json:"order_id" gorm:"not null"`
	ProductID uint       `json:"product_id" gorm:"not null"`
	Product   Product    `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int        `json:"quantity" gorm:"not null;default:0"`
	Price     int        `json:"price" gorm:"not null;default:0"`
	SubTotal  int        `json:"sub_total" gorm:"not null;default:0"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
