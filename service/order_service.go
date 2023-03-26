package service

import (
	"belajar-go-orm/models"
	"belajar-go-orm/repository"
	"context"
	"github.com/lucsky/cuid"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *models.CreateOrderRequest) (*models.Order, error)
}

type OrderServiceImpl struct {
	orderRepository       repository.OrderRepository
	orderDetailRepository repository.OrderDetailRepository
}

func (o OrderServiceImpl) CreateOrder(ctx context.Context, orderRequest *models.CreateOrderRequest) (*models.Order, error) {
	// begin transaction
	tx := o.orderRepository.BeginTransaction()

	totalPrice := 0

	for _, orderDetail := range orderRequest.Details {
		totalPrice += orderDetail.Price * orderDetail.Quantity
	}

	// create order here
	order := models.Order{
		OrderCode:  cuid.New(),
		TotalPrice: float64(totalPrice),
		UserID:     orderRequest.UserID,
	}

	orderResult, err := o.orderRepository.InsertOrder(ctx, tx, &order)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// create order detail here
	orderDetail := models.OrderDetail{
		OrderID:   order.ID,
		ProductID: 5,
		Quantity:  1,
		Price:     10000,
		SubTotal:  10000,
	}

	_, err = o.orderDetailRepository.InsertOrderDetail(ctx, tx, &orderDetail)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// commit transaction
	tx.Commit()

	return orderResult, nil
}

func NewOrderService(orderRepository repository.OrderRepository, orderDetailRepository repository.OrderDetailRepository) OrderService {
	return &OrderServiceImpl{orderRepository: orderRepository, orderDetailRepository: orderDetailRepository}
}
