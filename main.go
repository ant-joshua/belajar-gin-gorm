package main

import (
	"belajar-go-orm/controller"
	"belajar-go-orm/database"
	"belajar-go-orm/docs"
	"belajar-go-orm/models"
	"belajar-go-orm/repository"
	"belajar-go-orm/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucsky/cuid"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"time"
)

// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.DBConnection()
	if err != nil {
		fmt.Println(err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"

	orderRepository := repository.NewOrderRepository(db)
	orderDetailRepository := repository.NewOrderDetailRepository(db)

	orderService := service.NewOrderService(orderRepository, orderDetailRepository)
	//order := models.CreateOrderRequest{
	//	Details: []models.CreateOrderDetailRequest{
	//		{
	//			ProductID: 1,
	//			Quantity:  1,
	//			Price:     10000,
	//		},
	//		{
	//			ProductID: 2,
	//			Quantity:  1,
	//			Price:     20000,
	//		},
	//	},
	//	UserID: 1,
	//}
	//
	//orderResult, err := orderService.CreateOrder(context.Background(), &order)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("%+v", orderResult)

	// create gin server here
	startServer := gin.Default()
	v1 := startServer.Group("/api/v1")

	orderController := controller.NewOrderController(orderService)
	orderController.Routes(v1)

	startServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	startServer.Run(":" + os.Getenv("PORT"))
	fmt.Println("Server started at port " + os.Getenv("PORT"))

	//createUser("John Doe", 20)
	//getUserByID(1)
	//createProduct("Pro", 10000, 1)
	//updateUser(1, "John Doe Doe", 21)
	//createOrder(1, 2)
}

// create order repository with transaction
func createOrderRepository(userID uint, productID uint) {

}

func createOrder(userID uint, productID uint) {
	db := database.GetDB()

	generateCode := cuid.New()

	// begin a transaction
	tx := db.Debug().Begin()

	order := models.Order{
		OrderCode:  generateCode,
		TotalPrice: 0,
		UserID:     userID,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		//DeletedAt:  gorm.DeletedAt{},
	}
	err := tx.Debug().Create(&order).Error
	if err != nil {
		fmt.Println(err)
		//tx.Rollback()

	}

	// create order detail here
	orderDetail := models.OrderDetail{
		OrderID:   order.ID,
		ProductID: productID,
		Quantity:  1,
		Price:     10000,
		SubTotal:  10000,
	}

	err = tx.Debug().Create(&orderDetail).Error

	if err != nil {
		fmt.Println(err)
		//tx.Rollback()

	}

	// commit the transaction
	tx.Debug().Commit()

	return
}

func createProduct(name string, price float64, userID uint) {
	db := database.GetDB()
	product := models.Product{
		Name:   name,
		Price:  price,
		UserID: userID,
	}
	err := db.Create(&product).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func getUserByID(id int) {
	db := database.GetDB()
	var user models.User
	err := db.Preload("Product").First(&user, id).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}

func createUser(name string, age int) {
	db := database.GetDB()
	user := models.User{
		Name: name,
		Age:  age,
	}
	err := db.Debug().Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func updateUser(id int, name string, age int) {
	db := database.GetDB()
	var user models.User
	err := db.First(&user, id).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	user.Name = name
	user.Age = age
	err = db.Save(&user).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
