package model

import (
	"fmt"
	"log"
	"time"
)

type Business struct {
	ID           uint      `json:"id";gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"size:50"`
	Info         string    `json:"info"`
	RegisterTime time.Time `json:"register_time"`
}

func (Business) TableName() string {
	return "business"
}

// 根据商家 id 查询商家
// @param id uint "商家id"
// @param info string "商家信息"
// @return id uint "商家id"
func (b *Business) GetBusinessById() (business Business, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.GetBusinessById: %w", err)
		}
	}()
	result := DB.First(&business, b.ID)
	err = result.Error
	if err != nil {
		//log.Println(err)
		return Business{}, err
	}
	return
}

// 添加商家
func (b *Business) AddBusiness() (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.AddBusiness: %w", err)
		}
	}()
	result := DB.Create(b)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}
	return b.ID, err
}

// 删除商家
func (b *Business) DeleteBusiness() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.DeleteBusiness: %w", err)
		}
	}()
	result := DB.Delete(b)
	err = result.Error
	if err != nil {
		return err
	}
	return
}

// 更新商家
func (b *Business) UpdateBusiness(name, info string) (business Business, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.UpdateBusiness: %w", err)
		}
	}()
	// 判断数据库中是否有这个ID
	result := DB.First(&business, b.ID)
	err = result.Error
	if err != nil {
		return Business{}, err
	}
	// 更新数据
	result = DB.Model(b).Updates(Business{Name: name, Info: info})
	err = result.Error
	if err != nil {
		return Business{}, err
	}
	return *b, nil
}
