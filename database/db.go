package database

import (
	"belajar-go-orm/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

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
