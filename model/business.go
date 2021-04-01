package model

import (
	"time"
)

type Business struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"size:50"`
	Info         string    `json:"info"`
	ProductType  string    `json:"product_type" gorm:"size:2000"`
	RegisterTime time.Time `json:"register_time"`
}

func (Business)TableName() string {
	return "business"
}
