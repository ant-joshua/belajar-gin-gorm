package controller

import (
	"belajar-go-orm/models"
	"belajar-go-orm/service"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context)
	GetOrderById(ctx *gin.Context)
	Routes(r *gin.RouterGroup)
}

type OrderControllerImpl struct {
	orderService service.OrderService
}

// GetOrderById godoc
// @Summary Get Order By Id
// @Description Get Order By Id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order Id"
// @Success 200 {object} models.Order
// @Router /orders/{id} [get]
func (o OrderControllerImpl) GetOrderById(ctx *gin.Context) {

}

func (o OrderControllerImpl) Routes(r *gin.RouterGroup) {
	//TODO implement me
	r.POST("/orders", o.CreateOrder)
	r.GET("/orders/:id", o.GetOrderById)
}

// CreateOrder godoc
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.CreateOrderRequest true "Order"
// @Success 200 {object} models.Order
// @Router /orders [post]
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
