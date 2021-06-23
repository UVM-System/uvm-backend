package model

import (
	"fmt"
	"log"
)

type Machine struct {
	Id         uint     `json:"id" gorm:"primaryKey"`
	BusinessId uint     `json:"business_id" gorm:"not null" sql:"type:integer constraint fk_product_business REFERENCES business(id)"`
	Business   Business `json:"business" gorm:"ForeignKey:BusinessId;AssociationForeignKey:ID"`
	Location   string   `json:"location" gorm:"size:100;not null"`
	//Products          []Product `gorm:"ForeignKey:BusinessId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductCategories string `json:"product_categories" gorm:"size:2000;not null"`
	ModelId           uint   `json:"model_id" gorm:"not null"`
}

func (Machine) TableName() string {
	return "machine"
}

// 添加售货柜
func (m *Machine) AddMachine() (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.AddMachine: %w", err)
		}
	}()
	result := DB.Create(m)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}
	return m.Id, err
}
