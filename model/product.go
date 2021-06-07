package model

import (
	"fmt"
	"log"
	"time"
)

type Product struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Business   Business  `json:"business";gorm:"ForeignKey:BusinessId;AssociationForeignKey:ID"`
	BusinessId uint      `json:"business_id" gorm:"not null";sql:"type:integer constraint fk_product_business REFERENCES business(id)"`
	Name       string    `json:"name" gorm:"size:100; not null"`
	Info       string    `json:"info" gorm:"size:150; not null"`
	Number     int       `json:"number" gorm:"not null"`
	UpdateTime time.Time `json:"update_time"`
	Price      float64   `json:"price" gorm:"not null"`
	//Image 	   Image	 `json:"image" gorm:"ForeignKey:ImageID;AssociationForeignKey:ID"`
	//ImageID		uint	 `json:"image_id";gorm:"not null";sql:"type:integer constraint fk_product_image REFERENCES image(id)"`
}

func (Product) TableName() string {
	return "product"
}

// 根据商家ID查找所有商品
func GetProductList(businessId uint) (productList []Product, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.GetProductList: %w", err)
		}
	}()
	result := DB.Where(&Product{BusinessId: businessId}).Find(&productList)
	err = result.Error
	if err != nil {
		return nil, err
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
