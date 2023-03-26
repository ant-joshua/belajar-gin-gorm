package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Age       int       `json:"age" gorm:"not null;default:0"`
	Product   []Product `json:"product" gorm:"foreignKey:UserID"`
	Order     []Order   `json:"order" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
