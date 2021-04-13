package model

import (
	"log"
	"time"
)

type Business struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"size:50"`
	Info         string    `json:"info"`
	RegisterTime time.Time `json:"register_time"`
}

func (Business) TableName() string {
	return "business"
}

// 增加商家
// @param name string "商家名称"
// @param info string "商家信息"
// @return id uint "商家id"
func (*Business) AddBusiness(name, info string) (id uint, err error) {
	business := Business{
		Name:         name,
		Info:         info,
		RegisterTime: time.Now(),
	}
	result := DB.Create(&business)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}
	return business.ID, nil
}

// 根据商家 id 查询商家
// @param id uint "商家id"
// @param info string "商家信息"
// @return id uint "商家id"
func (*Business) GetBusinessById(id uint) (business Business, err error) {
	result := DB.First(&business, id)
	err = result.Error
	if err != nil {
		log.Println(err)
		return Business{}, result.Error
	}
	return
}
