package model

import "time"

type Order struct {
	Id uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id" gorm:"not null"`
	BusinessId uint `json:"business_id" gorm:"not null"`
	MachineId uint `json:"machine_id" gorm:"not null"`
	OrderNumber string `json:"order_number" gorm:"size:50; not null"`
	Status string `json:"status" gorm:"size:10; not null"`
	ProductAndNumber string `json:"product_and_number" gorm:"size:1000; not null"`
	Price int `json:"price" gorm:"default:0; not null"`
	CreatedTime time.Time `json:"created_time"`
}

func (Order)TableName() string {
	return "order"
}
