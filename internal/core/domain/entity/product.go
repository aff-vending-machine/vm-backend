package entity

import (
	"errors"
	"time"
)

type Product struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	SKU       string    `json:"sku" gorm:"uniqueIndex"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	ImageURL  string    `json:"image_url"`
	Price     float64   `json:"price"`
}

func (e Product) TableName() string {
	return "products"
}

func (e Product) Validate() error {
	if e.Name == "" {
		return errors.New("name is required")
	}
	if e.SKU == "" {
		return errors.New("sku is required")
	}
	if e.Price < 0 {
		return errors.New("price should be greater than or equal to zero")
	}
	return nil
}
