package controller

import (
	"belajar-go-orm/models"
	"belajar-go-orm/service"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context)
	Routes(r *gin.Engine)
}

type OrderControllerImpl struct {
	orderService service.OrderService
}

func (o OrderControllerImpl) Routes(r *gin.Engine) {
	//TODO implement me
	r.POST("/orders", o.CreateOrder)
}

func (o OrderControllerImpl) CreateOrder(ctx *gin.Context) {
	var request models.CreateOrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := o.orderService.CreateOrder(ctx, &request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, order)
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{orderService: orderService}
}
