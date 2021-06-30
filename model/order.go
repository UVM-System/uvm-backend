package model

import (
	"fmt"
	"log"
	"time"
)

type Order struct {
	Id           uint      `json:"id" gorm:"primaryKey"`
	UserId       uint      `json:"user_id"`
	BusinessId   uint      `json:"business_id"`
	MachineId    uint      `json:"machine_id"`
	OrderNumber  string    `json:"order_number" gorm:"size:50; not null"`
	Status       bool      `json:"status" gorm:"not null"`
	OrderContent string    `json:"order_content" gorm:"size:1000; not null"` // ProductId|Number|Price|Name
	Price        float64   `json:"price" gorm:"default:0; not null"`
	CreatedTime  time.Time `json:"created_time"`
}

func (Order) TableName() string {
	return "order"
}

/**
结构体查询得到OrderList
*/
func (o *Order) GetOrderListByStructQuery() (orders []Order, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.GetOrderListByStructQuery: %w", err)
		}
	}()
	result := DB.Where(o).Find(&orders)
	err = result.Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
func (o *Order) AddOrder() (Id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.AddOrder: %w", err)
		}
	}()
	result := DB.Create(o)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, err
	}
	return o.Id, nil
}
