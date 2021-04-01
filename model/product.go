package model

import "time"

type Product struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:100; not null"`
	Info string `json:"info" gorm:"size:150; not null"`
	Number int `json:"number" gorm:"not null"`
	UpdateTime time.Time `json:"update_time"`
}

func (Product)TableName() string {
	return "product"
}
