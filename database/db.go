package database

import (
	"belajar-go-orm/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=belajar_go_orm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection Failed to Open %v\n", err)
		panic(err)
		return db, err
	}

	db.Debug().AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderDetail{})

	fmt.Println("Connection Established")

	return db, nil
}

func GetDB() *gorm.DB {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	return db
}
