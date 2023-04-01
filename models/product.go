package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"type:varchar(100);not null" json:"name,omitempty"`
	UserID    uint       `json:"user_id" gorm:"not null" json:"userID,omitempty"`
	Price     float64    `json:"price" gorm:"not null;default:0.0" json:"price,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before Create product")

	if len(p.Name) < 4 {
		//err = tx.AddError(fmt.Errorf("name is required"))
		err = errors.New("name is too short (minimum is 4 characters)")
	}
	return
}
