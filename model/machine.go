package model

type Machine struct {
	Id                uint     `json:"id" gorm:"primaryKey"`
	BusinessId        uint     `json:"business_id" gorm:"not null";sql:"type:integer constraint fk_product_business REFERENCES business(id)"`
	Business          Business `json:"business";gorm:"ForeignKey:BusinessId;AssociationForeignKey:ID"`
	Location          string   `json:"location" gorm:"size:100;not null"`
	ProductCategories string   `json:"product_categories" gorm:"size:2000;not null"`
	ModelId           uint     `json:"model_id" gorm:"not null"`
}

func (Machine) TableName() string {
	return "machine"
}
