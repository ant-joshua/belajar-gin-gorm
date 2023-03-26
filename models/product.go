package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string  `json:"name" gorm:"type:varchar(100);not null"`
	UserID uint    `json:"user_id" gorm:"not null"`
	Price  float64 `json:"price" gorm:"not null;default:0.0"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before Create product")

	if len(p.Name) < 4 {
		//err = tx.AddError(fmt.Errorf("name is required"))
		err = errors.New("name is too short (minimum is 4 characters)")
	}
	return
}
