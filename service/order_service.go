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
	tx := o.orderRepository.BeginTransaction(ctx)

	totalPrice := 0
	orderDetails := make([]models.OrderDetail, 0)

	for _, orderDetail := range orderRequest.Details {
		totalPrice += orderDetail.Price * orderDetail.Quantity
	}

	// create order here
	order := models.Order{
		OrderCode:  cuid.New(),
		TotalPrice: float64(totalPrice),
		UserID:     orderRequest.UserID,
	}

	orderResult, err := o.orderRepository.InsertOrder(ctx, &order)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, orderDetailRequest := range orderRequest.Details {
		var orderDetail models.OrderDetail = models.OrderDetail{
			OrderID:   orderResult.ID,
			ProductID: orderDetailRequest.ProductID,
			Quantity:  orderDetailRequest.Quantity,
			Price:     orderDetailRequest.Price,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	for _, orderDetail := range orderDetails {
		_, err = o.orderDetailRepository.InsertOrderDetail(ctx, &orderDetail)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	orderResult.OrderDetail = orderDetails

	// commit transaction
	tx.Commit()

	return orderResult, nil
}

func NewOrderService(orderRepository repository.OrderRepository, orderDetailRepository repository.OrderDetailRepository) OrderService {
	return &OrderServiceImpl{orderRepository: orderRepository, orderDetailRepository: orderDetailRepository}
}
