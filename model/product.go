package model

import (
	"fmt"
	"log"
	"time"
)

type Product struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	BusinessId  uint      `json:"business_id"`
	Name        string    `json:"name" gorm:"size:100; not null"`
	EnglishName string    `json:"english_name" gorm:"size:100; not null"`
	Info        string    `json:"info" gorm:"size:150; not null"`
	Number      int       `json:"number" gorm:"not null"`
	UpdateTime  time.Time `json:"update_time"`
	Price       float64   `json:"price" gorm:"not null"`
	Image       Image     `json:"image" gorm:"ForeignKey:ImageID;AssociationForeignKey:ID"`
	ImageID     uint      `json:"image_id" gorm:"not null" sql:"type:integer constraint fk_product_image REFERENCES image(id)"`
}

//Business   Business  `json:"business" gorm:"ForeignKey:BusinessId;AssociationForeignKey:ID"`
//BusinessId uint      `json:"business_id" gorm:"not null";sql:"type:integer constraint fk_product_business REFERENCES business(id) on delete cascade on update cascade"`
// Product和Image是一对一关系；

func (Product) TableName() string {
	return "product"
}

/**
根据结构体查询条件查找一个商品（只适用于等号匹配）
*/
func (p *Product) GetProductByStructQuery() (product Product, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.GetProductByOneField: %w", err)
		}
	}()
	result := DB.Preload("Image").Where(p).First(&product)
	if result.Error != nil {
		log.Println(result.Error)
		return Product{}, result.Error
	}
	return
}

// 增加商品
func (p *Product) AddProduct() (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.AddProduct: %w", err)
		}
	}()
	result := DB.Create(p)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}
	return p.Id, err

}
func (p *Product) UpdateProduct() (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.UpdateProduct: %w", err)
		}
	}()
	// 判断数据库中是否有这个ID
	result := DB.First(&Product{}, p.Id)
	err = result.Error
	if err != nil {
		return 0, err
	}
	// 更新数据
	result = DB.Model(p).Updates(*p)
	err = result.Error
	if err != nil {
		return 0, err
	}
	return p.Id, nil
}
